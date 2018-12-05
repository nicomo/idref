// persname_t:(Natsume AND Sōseki)

package idref

import (
	"reflect"
	"testing"
)

type authSearchPersonTest struct {
	Description string
	Input       string
	Expected    Authorities
	Valid       bool
}

var authSearchPersonTestCases = []authSearchPersonTest{
	{
		Description: "Person Search: Natsume AND Sōseki",
		Input:       "Natsume Sōseki",
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
		Valid: false,
	},
}

func TestAuthSearchPerson(t *testing.T) {
	for _, test := range authSearchPersonTestCases {
		actual, err := AuthSearchPerson(test.Input)
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

func TestAuthSearchAll(t *testing.T) {
	for _, test := range authSearchPersonTestCases {
		actual, err := AuthSearchAll(test.Input)
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
