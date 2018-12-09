package idref

import (
	"reflect"
	"testing"
)

type idref2idTest struct {
	Description string
	Input       []string
	Expected    map[string][]Identifier
	Valid       bool
}

func TestIDRef2ID(t *testing.T) {

	idref2idTestCases := []idref2idTest{
		{
			Description: "Get sources IDs for valid PPN",
			Input:       []string{"139753753"},
			Expected: map[string][]Identifier{
				"139753753": []Identifier{
					{
						"238738325",
						"VIAF",
					},
					{
						"0000000385709539",
						"ISNI",
					},
					{
						"dacos",
						"HAL",
					},
					{
						"0000000293615295",
						"ORCID",
					},
					{
						"http://catalogue.bnf.fr/ark:/12148/cb16180961h",
						"BNF",
					},
				},
			},
			Valid: true,
		},
	}

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
