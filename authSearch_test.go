package idref

import (
	"reflect"
	"testing"
)

type authSearchTest struct {
	Description      string
	InputSearchTerms string
	InputSearchIndex string
	Expected         Authorities
	Valid            bool
}

var authSearchTestCases = []authSearchTest{
	{
		Description:      "Person Search, one result",
		InputSearchTerms: "Natsume Sōseki",
		InputSearchIndex: "persname_t",
		Expected: Authorities{
			AuthorityRecord{
				ID: "027044971",
				Person: Person{
					PrefLabel: "Natsume, Sōseki (1867-1916)",
					AltLabels: []string{
						"Natsume, Sōseki (1867-1916)",
						"夏目, 漱石 (1867-1916)",
						"Sōseki",
						"Natsume Sōseki",
						"Natsume, Kin'nosuke",
						"Natsume Kin'nosuke",
						"Kin'nosuke Natsume",
						"Natsume Kinnosuke",
						"Kinnosuke Natsume",
						"Sōseki, Natsume",
						"Xiamu, Shushi",
					},
				},
			},
		},
		Valid: true,
	},
	{
		Description:      "Org Search, one result",
		InputSearchTerms: "Lycée Michelet",
		InputSearchIndex: "corpname_t",
		Expected:         Authorities{},
		Valid:            false,
	},
}

func TestAuthSearch(t *testing.T) {
	for _, test := range authSearchTestCases {
		actual, err := AuthSearch(test.InputSearchTerms, test.InputSearchIndex)
		if err != nil && !test.Valid {
			t.Logf("PASS %s: got %v", test.Description, err)
		}
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS %s", test.Description)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.Description, test.Expected, actual)
		}
	}
}
