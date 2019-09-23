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
		Description:      "Person Search, valid",
		InputSearchTerms: "Natsume Sōseki",
		InputSearchIndex: "persname_t",
		Expected: Authorities{
			AuthorityRecord{
				ID: "027044971",
				Identifiers: []Identifier{
					Identifier{
						ID:     "027044971",
						Source: "idref",
					},
				},
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
		Description:      "Person Search, test",
		InputSearchTerms: "Roberto WOLFLER-CALVO",
		InputSearchIndex: "persname_t",
		Expected: Authorities{
			AuthorityRecord{
				ID: "130072192",
				Identifiers: []Identifier{
					Identifier{
						ID:     "130072192",
						Source: "idref",
					},
				},
				Person: Person{
					PrefLabel: "Wolfler-Calvo, Roberto (1964-....)",
					AltLabels: []string{
						"Wolfler-Calvo, Roberto (1964-....)",
					},
				},
			},
		},
		Valid: true,
	},
	{
		Description:      "Org Search, valid",
		InputSearchTerms: "lamentin aéroport",
		InputSearchIndex: "corpname_t",
		Expected: Authorities{
			AuthorityRecord{
				ID: "027767892",
				Identifiers: []Identifier{
					Identifier{
						ID:     "027767892",
						Source: "idref",
					},
				},
				Organization: Organization{
					PrefLabel: "Aéroport du Lamentin (Fort-de-France)",
					AltLabels: []string{
						"Aéroport du Lamentin (Fort-de-France)",
						"Le Lamentin (Martinique). Aéroport",
						"Fort-de-France (Martinique) -- Aéroport du Lamentin",
					},
				},
			},
			AuthorityRecord{
				ID: "026357968",
				Identifiers: []Identifier{
					Identifier{
						ID:     "026357968",
						Source: "idref",
					},
				},
				Organization: Organization{
					PrefLabel: "Aéroport international de Fort-de-France-Le Lamentin",
					AltLabels: []string{
						"Aéroport international de Fort-de-France-Le Lamentin",
						"Aéroport de Fort-de-France-Le Lamentin",
						"Aéroport Le Lamentin (Fort-de-France)",
						"Chambre de commerce et d'industrie (Fort-de-France). Aéroport international de Fort-de-France-Le Lamentin",
					},
				},
			},
		},
		Valid: true,
	},
	{
		Description:      "Person Search with char parsing of name",
		InputSearchTerms: "Natsume K Sōseki",
		InputSearchIndex: "persname_t",
		Expected: Authorities{
			AuthorityRecord{
				ID: "027044971",
				Identifiers: []Identifier{
					Identifier{
						ID:     "027044971",
						Source: "idref",
					},
				},
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
			t.Fatalf("FAIL for %s: expected %+v, actual result was %+v", test.Description, test.Expected, actual)
		}
	}
}
