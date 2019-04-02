package idref

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/beevik/etree"
)

// AuthoritiesGet with []string, returns map[string][]AuthorityRecord
// FIXME: we're burrying the error in the AuthorityResult struct,
// which might not be the best way to proceed
func AuthoritiesGet(PPNS []string) []AuthorityResult {
	jobs := make(chan string, len(PPNS))
	results := make(chan AuthorityResult, len(PPNS))

	// dispatch jobs to number of workers, capping at 5
	numW := 5
	if len(PPNS) < 5 {
		numW = len(PPNS)
	}
	// This starts up to 5 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= numW; w++ {
		go worker(w, jobs, results)
	}

	// Here we send the `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for _, v := range PPNS {
		jobs <- v
	}
	close(jobs)

	// Finally we collect all the results of the work
	var res []AuthorityResult
	for a := 1; a <= numW; a++ {
		res = append(res, <-results)
	}
	return res
}

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`.
func worker(id int, jobs <-chan string, results chan<- AuthorityResult) {
	for j := range jobs {
		record, err := AuthorityGet(j)
		if err != nil {
			results <- AuthorityResult{
				Err: err,
			}
		}

		results <- AuthorityResult{
			Auth: record,
		}
	}
}

// AuthorityGet retrieves an authority record
// given an ID in the IdRef databases
func AuthorityGet(PPN string) (AuthorityRecord, error) {

	// provision an authority struct
	auth := AuthorityRecord{}

	getURL := "https://www.idref.fr/" + PPN + ".rdf"
	resp, err := callIDRef(getURL)
	if err != nil {
		return auth, fmt.Errorf("couldn't retrieve response from IdRef: %v", err)
	}

	result := etree.NewDocument()
	if err = result.ReadFromBytes(resp); err != nil {
		return auth, err
	}

	// loop through the documents until we find the
	// document with the authority's metadata
	for _, e := range result.FindElements("./RDF/bibo:Document") {
		about := e.SelectAttrValue("about", "unknown")
		if strings.HasSuffix(about, PPN) {
			// this is a special Document
			// namely the header for this very authority record
			for _, t := range e.ChildElements() {
				switch headerTag := t.Tag; headerTag {
				case "created":
					auth.DateCreated = t.Text()
				case "identifier":
					auth.ID = t.Text()
					auth.Identifiers = append(auth.Identifiers,
						Identifier{
							ID:     t.Text(),
							Source: "idref",
						},
					)
				case "modified":
					auth.DateUpdated = t.Text()
				}
			}
			break
		}
	}

	// Is this the authority for a Person?
	if p := result.FindElement("RDF/foaf:Person"); p != nil {
		person := Person{}
		for _, e := range p.ChildElements() {
			switch tag := e.Tag; tag {
			case "prefLabel":
				person.PrefLabel = e.Text()
			case "name":
				person.Name = e.Text()
			case "familyName":
				person.FamilyName = e.Text()
			case "givenName":
				person.GivenName = e.Text()
			case "surname":
				person.Surname = e.Text()
			case "citizenship":
				s := e.SelectAttrValue("resource", "dflt")
				if s != "dflt" && len(s) > 3 {
					person.Citizenship = s
				}
			case "language":
				s := e.SelectAttrValue("resource", "dflt")
				if s != "dflt" && len(s) > 3 {
					person.Language = s[len(s)-3:]
				}
			case "orcidId":
				auth.Identifiers = append(
					auth.Identifiers,
					Identifier{
						ID:     e.Text(),
						Source: "ORCID",
					},
				)
			case "FRBNF":
				auth.Identifiers = append(
					auth.Identifiers,
					Identifier{
						ID:     e.Text(),
						Source: "FRBNF",
					},
				)
			case "identifierValid":
				auth.Identifiers = append(
					auth.Identifiers,
					Identifier{
						ID:     e.Text(),
						Source: "ISNI",
					},
				)
			default:
				continue
			}
		}

		if birth := p.FindElement("bio:event/bio:Birth/bio:date"); birth != nil {
			person.DateBirth = birth.Text()
		}
		if death := p.FindElement("bio:event/bio:Death/bio:date"); death != nil {
			person.DateDeath = death.Text()
		}

		auth.Person = person
	}

	// Is this the authority for an Organization?
	if o := result.FindElement("RDF/foaf:Organization"); o != nil {
		org := Organization{}
		for _, e := range o.ChildElements() {
			switch tag := e.Tag; tag {
			case "prefLabel":
				org.PrefLabel = e.Text()
			case "name":
				org.Name = e.Text()
			case "altLabel":
				org.AltLabels = append(org.AltLabels, e.Text())
			case "dateOfBirth":
				org.DateOfBirth = e.Text()
			case "FRBNF":
				auth.Identifiers = append(
					auth.Identifiers,
					Identifier{
						ID:     e.Text(),
						Source: "FRBNF",
					},
				)
			case "identifierValid":
				auth.Identifiers = append(
					auth.Identifiers,
					Identifier{
						ID:     e.Text(),
						Source: "ISNI",
					},
				)
			case "rnsr":
				auth.Identifiers = append(
					auth.Identifiers,
					Identifier{
						ID:     e.Text(),
						Source: "RNSR",
					},
				)
			default:
				continue
			}
		}
		auth.Organization = org
	}

	return auth, nil

	/* TODO: other types of auth records I'm not interested in right now:
	Nom géographique
	Nom générique de famille
	Titre uniforme
	Auteur / Titre
	Nom commun RAMEAU
	Nom commun FMeSH
	*/
}

// callIDRef performs the http GET
// retrieves the response and puts it in a slice of bytes
func callIDRef(getURL string) ([]byte, error) {
	// get the result from the url
	resp, err := http.Get(getURL)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	// check http errors and anything other than 200
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got %s", resp.Status)
	}

	// put the response into a []byte
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}
