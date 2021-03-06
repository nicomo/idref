package idref

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"unicode"

	"github.com/beevik/etree"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// AuthSearch uses the Solr search at IdRef
// Can retrieve Persons, Organization, etc
// defaults to the "all" index if the provided index is unknown or not implemented
func AuthSearch(s, index string) (Authorities, error) {
	// we provision a slice of authorities
	auths := Authorities{}

	// build search string
	searchString, err := buildSearchString(s)
	if err != nil {
		return auths, err
	}

	index = validateIndex(index)

	// build URL
	qURLString, err := qURLBuild(searchString, index)
	if err != nil {
		log.Fatal(err)
	}

	// actually call the web service
	resp, err := callIDRef(qURLString)
	if err != nil {
		return auths, fmt.Errorf("could not retrieve response from IdRef: %v - %s", err, qURLString)
	}

	result := etree.NewDocument()
	if err = result.ReadFromBytes(resp); err != nil {
		return auths, err
	}

	for _, doc := range result.FindElements("./response/result/doc") {
		auth := AuthorityRecord{}
		// what sort of authority is this?
		if rTZ := doc.FindElement("[@name='recordtype_z']"); rTZ != nil {
			switch rTZ.Text() {
			case "a":
				parsePerson(doc, &auth)
				auths = append(auths, auth)
			case "b":
				parseOrg(doc, &auth)
				auths = append(auths, auth)
			default:
				fmt.Println("recordtype_z unknown or not implemented")
			}
		}
	}

	return auths, nil
}

// buildSearchString removes words < 2 and diacritics
// they're not supported by idref search solr :-(
func buildSearchString(s string) (string, error) {

	// we stop if the whole search string is < 3 chars
	if len(s) < 3 {
		return "", fmt.Errorf("search string has to be at least 3 char long, is %d", len(s))
	}

	s1 := strings.Split(s, " ")

	// remove strings <2 chars
	s2 := []string{}
	for _, v := range s1 {
		if len(v) >= 2 {
			s2 = append(s2, v)
		}
	}

	// remove diacritics
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

	for i, word := range s2 {
		b := make([]byte, len(word))
		_, _, e := t.Transform(b, []byte(word), true)
		if e != nil {
			log.Fatalln(e)
		}
		s2[i] = string(b)
	}

	return strings.Join(s2, " AND "), nil

}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// parsePerson parses an xml tree into a Person auth struct
func parsePerson(doc *etree.Element, auth *AuthorityRecord) {

	if arr := doc.FindElement("arr[@name='affcourt_r']"); arr != nil {
		for _, strTag := range arr.SelectElements("str") {
			auth.Person.AltLabels = append(auth.Person.AltLabels, strTag.Text())
		}
	}

	for _, v := range doc.FindElements("str") {
		for _, attr := range v.Attr {
			switch attr.Value {
			case "ppn_z":
				auth.ID = v.Text()
				auth.Identifiers = append(
					auth.Identifiers,
					Identifier{
						ID:     v.Text(),
						Source: "idref",
					},
				)
			case "affcourt_z":
				auth.Person.PrefLabel = v.Text()
			}
		}
	}
}

func parseOrg(doc *etree.Element, auth *AuthorityRecord) {

	if arr := doc.FindElement("arr[@name='affcourt_r']"); arr != nil {
		for _, strTag := range arr.SelectElements("str") {
			auth.Organization.AltLabels = append(auth.Organization.AltLabels, strTag.Text())
		}
	}

	for _, v := range doc.FindElements("str") {
		for _, attr := range v.Attr {
			switch attr.Value {
			case "ppn_z":
				auth.ID = v.Text()
				auth.Identifiers = append(
					auth.Identifiers,
					Identifier{
						ID:     v.Text(),
						Source: "idref",
					},
				)
			case "affcourt_z":
				auth.Organization.PrefLabel = v.Text()
			}
		}
	}

	for _, v := range doc.FindElements("date") {
		for _, attr := range v.Attr {
			switch attr.Value {
			case "datenaissance_dt":
				auth.Organization.DateOfBirth = v.Text()
			case "dateetat_dt":
				auth.DateCreated = v.Text()
			}
		}
	}

}

// qURLBuild builds the url search query
func qURLBuild(searchString, index string) (string, error) {

	qURL, err := url.Parse("https://www.idref.fr/Sru/Solr")
	if err != nil {
		return "", fmt.Errorf("couldn't create URL: %s", err)
	}
	q := qURL.Query()
	q.Set("q", index+":("+searchString+")")
	q.Add("start", "0")
	q.Add("rows", "30")
	q.Add("fl", "recordtype_z,ppn_z,affcourt_r,affcourt_z")
	qURL.RawQuery = q.Encode()

	return qURL.String(), nil
}

// validateIndex defaults to "all" if the index is unknown
func validateIndex(index string) string {
	indexes := []string{"persname_t", "persname_s", "corpname_t", "corpname_s"}
	for _, v := range indexes {
		if index == v {
			return index
		}
	}
	return "all"
}
