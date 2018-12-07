package idref

import (
	"fmt"

	"github.com/beevik/etree"
)

//
// ID2IDRef takes a single ID from another source (VIAF, BNF, etc)
// and retrieves the IdRef ID (PPN)
func ID2IDRef(id string) (string, error) {
	getURL := "https://www.idref.fr/services/id2idref/" + id

	resp, err := callIDRef(getURL)
	if err != nil {
		return "", fmt.Errorf("couldn't retrieve response from IdRef: %v", err)
	}

	result := etree.NewDocument()
	if err := result.ReadFromBytes(resp); err != nil {
		return "", err
	}

	if ppn := result.FindElement("//ppn"); ppn != nil {
		return ppn.Text(), nil
	}

	return "", fmt.Errorf("could not find a PPN")
}

// TODO: func IDs2IDRef with []string
