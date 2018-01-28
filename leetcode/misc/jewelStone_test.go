package misc

import (
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	cases := []struct {
		JewelInput  string
		StonesInput string
		Expected    int
	}{
		{"", "AsdhkkD", 0},
		{"Ll", "", 0},
		{"a", "aAaAa", 3},
		{"be", "BebE", 2},
	}
	for _, c := range cases {
		actual := NumJewelsInStones(c.JewelInput, c.StonesInput)
		if actual != c.Expected {
			t.Fatalf("JewelInput: %s, StonesInput: %s, output: %d, expected: %d", c.JewelInput, c.StonesInput, actual, c.Expected)
		}
	}
}
