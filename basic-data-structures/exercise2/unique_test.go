package exercise2

import (
	"slices"
	"strconv"
	"testing"
)

type testcase struct {
	strings  []string
	expected []string
}

var testcases = []testcase{
	{[]string{"a", "b", "a", "b", "c", "d", "a"}, []string{"a", "b", "c", "d"}},
	{[]string{"9", "1", "1", "1", "123", "123", "9"}, []string{"9", "1", "123"}},
	{[]string{"9", "", "", "1", "123", "112", "9"}, []string{"9", "", "1", "123", "112"}},
}

func TestRemoveDuplicates(t *testing.T) {
	for i, c := range testcases {
		t.Run("testcase_"+strconv.Itoa(i+1), func(t *testing.T) {
			got := RemoveDuplicates(c.strings)

			if !slices.Equal(c.expected, got) {
				t.Fatalf("strings: %v, expected: %v, got: %v", c.strings, c.expected, got)
			}
		})
	}
}
