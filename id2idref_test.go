package idref

import (
	"reflect"
	"testing"
)

type id2idrefTest struct {
	Description string
	Input       string
	Expected    string
	Valid       bool
}

func TestID2IDRef(t *testing.T) {

	id2idrefTestCases := []id2idrefTest{
		{
			Description: "Get IdRef ID from other source ID",
			Input:       "Q12328808",
			Expected:    "113592051",
			Valid:       true,
		},
		{
			Description: "No IdRef ID",
			Input:       "Q123",
			Expected:    "",
			Valid:       false,
		},
	}

	for _, test := range id2idrefTestCases {
		actual, err := ID2IDRef(test.Input)
		if err != nil && !test.Valid {
			t.Logf("PASS %s: got %v", test.Description, err)
			continue
		}
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS %s", test.Description)
		} else {
			t.Fatalf("FAIL for %s (%s): \nexpected %v\nactual result was %v", test.Input, test.Description, test.Expected, actual)
		}
	}
}
