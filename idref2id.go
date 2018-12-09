package idref

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
)

// IDRef2ID takes one or more IdREf PPNs (IDs) and retrieves
// IDs for the same entity in other sources (VIAF, BNF, etc)
func IDRef2ID(ppns []string) (map[string][]Identifier, error) {

	qPPNs := strings.Join(ppns, ",")
	getURL := "https://www.idref.fr/services/idref2id/" + qPPNs

	resp, err := callIDRef(getURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve response from IdRef: %v", err)
	}

	result := etree.NewDocument()
	if err := result.ReadFromBytes(resp); err != nil {
		return nil, err
	}

	// provision a result map with query ppns as keys
	m := make(map[string][]Identifier)

	// for each query...
	for _, query := range result.FindElements("./sudoc/query") {
		// we get the ppn (IdRef ID) and corresponding source + identifier at source
		qPPN := query.FindElement("ppn")
		source := query.FindElement("result/source")
		id := query.FindElement("result/identifiant")

		if qPPN != nil && source != nil && id != nil && source.Text() != "" {
			// use query PPN as key and add Identifier as value in result slice
			m[qPPN.Text()] = append(m[qPPN.Text()], Identifier{
				Source: source.Text(),
				ID:     id.Text(),
			})
		}

	}

	return m, nil
}
