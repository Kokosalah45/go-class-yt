package hello

import "testing"

type testCases struct {
	items []string
	result string
}

func TestSay(t *testing.T) {
	testCases := []testCases{
		{[]string{"world"}, "Hello, world"},
		{[]string{"world", "mars"}, "Hello, world, mars"},
		{[]string{}, "Hello, World"},
	}

	for _,tc := range testCases {
		actual := Say(tc.items)
		if actual != tc.result {
			t.Errorf("Expected: %v, but got: %v", tc.result, actual)
		}
	}
}
