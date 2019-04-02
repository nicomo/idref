package idref

import (
	"reflect"
	"testing"
)

type authGetSingleTest struct {
	Description string
	Input       string
	Expected    AuthorityRecord
	Valid       bool
}

type authGetSliceTest struct {
	Description string
	Input       []string
	Expected    []AuthorityResult
	Valid       bool
}

func TestAuthorityGet(t *testing.T) {
	authGetSingleTestCases := []authGetSingleTest{
		{
			Description: "Single PPN, valid Org: Lycée Michelet",
			Input:       "035340096",
			Expected: AuthorityRecord{
				ID:          "035340096",
				DateCreated: "1998-02-06",
				DateUpdated: "2015-10-02T04:31:15",
				Identifiers: []Identifier{
					Identifier{
						ID:     "13182832",
						Source: "FRBNF",
					},
					Identifier{
						ID:     "12804817",
						Source: "FRBNF",
					},
				},
				Person: Person{},
				Organization: Organization{
					PrefLabel: "Lycée Michelet (Vanves, Hauts-de-Seine)",
					Name:      "Lycée Michelet (Vanves, Hauts-de-Seine)",
					AltLabels: []string{
						"Lycée de Vanves",
						"Petit collège de Vanves",
						"Lycée du Prince impérial",
					},
				},
			},
			Valid: true,
		},
		{
			Description: "Single PPN, valid Person: Pierre Bourdieu",
			Input:       "027715078",
			Expected: AuthorityRecord{
				ID:          "027715078",
				DateCreated: "1985-03-18",
				DateUpdated: "2018-11-20T15:12:51",
				Identifiers: []Identifier{
					Identifier{
						ID:     "11893402",
						Source: "FRBNF",
					},
					Identifier{
						ID:     "0000000121385892",
						Source: "ISNI",
					},
				},
				Person: Person{
					DateBirth:   "1930-08-01",
					DateDeath:   "2002-01-23",
					FamilyName:  "Bourdieu",
					GivenName:   "Pierre",
					Name:        "Bourdieu, Pierre",
					PrefLabel:   "Bourdieu, Pierre (1930-2002)",
					Language:    "fra",
					Citizenship: "http://sws.geonames.org/3017382",
				},
			},
			Valid: true,
		},
		{
			Description: "Single PPN, valid Person: JM Bonnisseau",
			Input:       "033522448",
			Expected: AuthorityRecord{
				ID:          "033522448",
				DateCreated: "1995-01-04",
				DateUpdated: "2019-04-01T18:01:34",
				Identifiers: []Identifier{
					Identifier{
						ID:     "12438130",
						Source: "FRBNF",
					},
					Identifier{
						ID:     "0000000281562229",
						Source: "ORCID",
					},
				},
				Person: Person{
					FamilyName:  "Bonnisseau",
					GivenName:   "Jean-Marc",
					Name:        "Bonnisseau, Jean-Marc",
					PrefLabel:   "Bonnisseau, Jean-Marc (1957-....)",
					Language:    "fra",
					Citizenship: "http://sws.geonames.org/3017382",
				},
			},
			Valid: true,
		},
		{
			Description: "Invalid PPN",
			Input:       "03534009",
			Expected:    AuthorityRecord{},
			Valid:       false,
		},
	}

	for _, test := range authGetSingleTestCases {
		actual, err := AuthorityGet(test.Input)
		if err != nil && !test.Valid {
			t.Logf("PASS %s: got %v", test.Description, err)
		}
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS %s", test.Description)
		} else {
			t.Fatalf("FAIL for %s (%s): expected %+v, actual result was %+v", test.Input, test.Description, test.Expected, actual)
		}
	}

}

func TestAuthoritiesGet(t *testing.T) {

	authGetSliceTestCases := []authGetSliceTest{
		{
			Description: "valid slice of ppns",
			Input:       []string{"035340096", "027715078"},
			Expected: []AuthorityResult{
				{
					Auth: AuthorityRecord{
						ID:          "035340096",
						DateCreated: "1998-02-06",
						DateUpdated: "2015-10-02T04:31:15",
						Identifiers: []Identifier{
							Identifier{
								ID:     "13182832",
								Source: "FRBNF",
							},
							Identifier{
								ID:     "12804817",
								Source: "FRBNF",
							},
						},
						Organization: Organization{
							PrefLabel: "Lycée Michelet (Vanves, Hauts-de-Seine)",
							Name:      "Lycée Michelet (Vanves, Hauts-de-Seine)",
							AltLabels: []string{
								"Lycée de Vanves",
								"Petit collège de Vanves",
								"Lycée du Prince impérial",
							},
						},
					},
				},
				{
					Auth: AuthorityRecord{
						ID:          "027715078",
						DateCreated: "1985-03-18",
						DateUpdated: "2018-11-20T15:12:51",
						Identifiers: []Identifier{
							Identifier{
								ID:     "11893402",
								Source: "FRBNF",
							},
							Identifier{
								ID:     "0000000121385892",
								Source: "ISNI",
							},
						},
						Person: Person{
							DateBirth:  "1930-08-01",
							DateDeath:  "2002-01-23",
							FamilyName: "Bourdieu",
							GivenName:  "Pierre",
							Name:       "Bourdieu, Pierre",
							PrefLabel:  "Bourdieu, Pierre (1930-2002)",
						},
					},
				},
			},
			Valid: true,
		},
	}
	for _, test := range authGetSliceTestCases {
		actual := AuthoritiesGet(test.Input)
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS %s", test.Description)
		} else {
			t.Fatalf("FAIL for %s (%s): expected %v, actual result was %v", test.Input, test.Description, test.Expected, actual)
		}
	}

}
