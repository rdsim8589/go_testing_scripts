package misc

import (
	"testing"
)

func TestIsIdealPermutation(t *testing.T) {
	cases := []struct {
		Input    []int
		Expected bool
	}{
		{[]int{1, 2, 0}, false},
		{[]int{1, 0, 2}, true},
	}
	for _, c := range cases {
		actual := IsIdealPermutation(c.Input)
		if actual != c.Expected {
			t.Fatalf("input: %v, output: %t, expected: %t", c.Input, actual, c.Expected)
		}
	}
}
