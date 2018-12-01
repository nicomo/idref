package idref

import (
	"fmt"

	"github.com/beevik/etree"
)

// DocsGet retrieve the documents associated
// with a single authority record
// the IdRef ID (a.k.a PPN) of the auth record should be provided
func DocsGet(PPN string) (Documents, error) {

	getURL := "https://www.idref.fr/services/references/" + PPN + ".xml"
	resp, err := callIDRef(getURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve response from IdRef: %v", err)
	}

	result := etree.NewDocument()
	if err := result.ReadFromBytes(resp); err != nil {
		return nil, err
	}

	// we provision a slice of docs
	docs := Documents{}

	// for each role...
	for _, role := range result.FindElements("./sudoc/result/role") {
		ar := AuthorityRole{
			UnimarcCode: role.SelectElement("unimarcCode").Text(),
			Marc21Code:  role.SelectElement("marc21Code").Text(),
			RoleName:    role.SelectElement("roleName").Text(),
		}
		// ... we parse the doc...
		for _, doc := range role.SelectElements("doc") {
			d := Document{
				AuthorityRole: ar,
				Citation:      doc.SelectElement("citation").Text(),
				ID:            doc.SelectElement("id").Text(),
				Source:        doc.SelectElement("referentiel").Text(),
			}
			// ... and add it to the docs
			docs = append(docs, d)
		}
	}

	return docs, nil
}
