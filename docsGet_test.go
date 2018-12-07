package idref

import (
	"reflect"
	"testing"
)

type docsGetTest struct {
	Description string
	Input       string
	Expected    Documents
	Valid       bool
}

func TestDocsGet(t *testing.T) {
	var docsGetTestCases = []docsGetTest{
		{
			Description: "Docs for single PPN, valid Person: Illiano",
			Input:       "197281397",
			Expected: Documents{
				{AuthorityRole: AuthorityRole{
					UnimarcCode: "070",
					Marc21Code:  "aut",
					RoleName:    "Auteur",
				},
					Citation: "Amélioration de la performance d'une ligne de production de vaccins grâce aux outils du lean manufacturing  / Marion Illiano  ; sous la direction de Philippe Lawton / [Lieu de publication inconnu] : [éditeur inconnu] , 2016",
					ID:       "19733184X",
					Source:   "sudoc",
					URI:      "",
					URL:      "",
				},
			},
			Valid: true,
		},
		{
			Description: "Docs for single invalid PPN",
			Input:       "02797524",
			Expected:    nil,
			Valid:       false,
		},
	}
	for _, test := range docsGetTestCases {
		actual, err := DocsGet(test.Input)
		if err != nil && !test.Valid {
			t.Logf("PASS %s: got %v", test.Description, err)
		}
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS %s", test.Description)
		} else {
			t.Fatalf("FAIL for %s (%s): expected %v, actual result was %v", test.Input, test.Description, test.Expected, actual)
		}
	}

}
