package idref

import (
	"fmt"

	"github.com/beevik/etree"
)

// IDRef2ID takes a signe IdREf PPN and retrieves
// IDs for the same entity in other sources (VIAF, BNF, etc)
func IDRef2ID(ppn string) ([]Identifier, error) {
	getURL := "https://www.idref.fr/services/idref2id/" + ppn

	resp, err := callIDRef(getURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve response from IdRef: %v", err)
	}

	result := etree.NewDocument()
	if err := result.ReadFromBytes(resp); err != nil {
		return nil, err
	}

	// provision slice of Identifier
	var Identifiers []Identifier

	// for each result...
	for _, res := range result.FindElements("./sudoc/query/result") {
		if source := res.SelectElement("source"); source != nil {
			if id := res.SelectElement("identifiant"); id != nil && source.Text() != "" {
				Identifiers = append(Identifiers, Identifier{
					Source: source.Text(),
					ID:     id.Text(),
				})
			}
		}
	}

	return Identifiers, nil
}

// TODO: func IDRefs2ID with []string
