package idref

import (
	"reflect"
	"testing"
)

type idref2idTest struct {
	Description string
	Input       string
	Expected    []Identifier
	Valid       bool
}

var idref2idTestCases = []idref2idTest{
	{
		Description: "Get sources IDs for valid PPN",
		Input:       "139753753",
		Expected: []Identifier{
			Identifier{
				"238738325",
				"VIAF",
			},
			Identifier{
				"0000000385709539",
				"ISNI",
			},
			Identifier{
				"dacos",
				"HAL",
			},
			Identifier{
				"0000000293615295",
				"ORCID",
			},
			Identifier{
				"http://catalogue.bnf.fr/ark:/12148/cb16180961h",
				"BNF",
			},
		},
		Valid: true,
	},
}

func TestIDRef2ID(t *testing.T) {
	for _, test := range idref2idTestCases {
		actual, err := IDRef2ID(test.Input)
		if err != nil && !test.Valid {
			t.Logf("PASS %s: got %v", test.Description, err)
		}
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS %s", test.Description)
		} else {
			t.Fatalf("FAIL for %s (%s): \nexpected %v\nactual result was %v", test.Input, test.Description, test.Expected, actual)
		}
	}

}
