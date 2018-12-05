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

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// AuthSearchPerson searches for a Person Authority
func AuthSearchPerson(s string) (Authorities, error) {

	// get search terms from request
	// remove words < 2
	// remove diacritics - they're not supported by idref search solr :-(
	sTerms := strings.Split(s, " ")
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

	for i, w := range sTerms {
		if len(w) < 2 {
			sTerms = append(sTerms[:i], sTerms[i+1:]...)
			continue
		}
		b := make([]byte, len(w))
		_, _, e := t.Transform(b, []byte(w), true)
		if e != nil {
			log.Fatalln(e)
		}
		sTerms[i] = string(b)
	}

	searchString := strings.Join(sTerms, " AND ")

	// build URL
	qURLString, err := qURLBuild(searchString)
	if err != nil {
		log.Fatal(err)
	}

	// we provision a slice of authorities
	auths := Authorities{}

	// actually call the web service
	resp, err := callIDRef(qURLString)
	if err != nil {
		return auths, fmt.Errorf("couldn't retrieve response from IdRef: %v", err)
	}

	result := etree.NewDocument()
	if err = result.ReadFromBytes(resp); err != nil {
		return auths, err
	}
	for _, doc := range result.FindElements("./response/result/doc") {

		auth := AuthorityRecord{}

		// what sort of authority is this?
		if rTZ := doc.FindElement("[@name='recordtype_z']"); t != nil {
			switch rTZ.Text() {
			case "a":
				parsePerson(doc, &auth)
			default:
				fmt.Println("recordtype_z unknown or not implemented)")
			}
		}

		auths = append(auths, auth)
	}

	return auths, nil
}

// qURLBuild builds the url search query
func qURLBuild(searchString string) (string, error) {

	qURL, err := url.Parse("https://www.idref.fr/Sru/Solr")
	if err != nil {
		return "", fmt.Errorf("couldn't create URL: %s", err)
	}
	q := qURL.Query()
	q.Set("q", "persname_t:("+searchString+")")
	q.Add("start", "0")
	q.Add("rows", "30")
	q.Add("fl", "recordtype_z,ppn_z,affcourt_r,affcourt_z")
	qURL.RawQuery = q.Encode()

	return qURL.String(), nil
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
			case "affcourt_z":
				auth.Person.PrefLabel = v.Text()
			}
		}
	}
}
