package exercise1

import (
	"strconv"
	"testing"
)

type testcase struct {
	id       int
	str      string
	expected bool
}

var testcases = []testcase{
	{1, "22/2/22", true},
	{2, "abab", false},
	{3, "121", true},
	{4, "d", true},
	{5, "ab", false},
	{6, "abc", false},
	{7, "", false},
	{8, "abc cba abc cba", true},
	{9, "Mr. Owl ate my metal worm", true},
}

func BenchmarkIsPalindrome1(b *testing.B) {
	for _, c := range testcases {
		b.Run("testcase"+strconv.Itoa(c.id), func(b *testing.B) {
			got := IsPalindrome1(c.str)

			if c.expected != got {
				b.Fatalf("str: %s, expected: %v, got: %v", c.str, c.expected, got)
			}
		})
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for _, c := range testcases {
		b.Run("testcase"+strconv.Itoa(c.id), func(b *testing.B) {
			got := IsPalindrome2(c.str)

			if c.expected != got {
				b.Fatalf("str: %s, expected: %v, got: %v", c.str, c.expected, got)
			}
		})
	}
}

func BenchmarkIsPalindrome3(b *testing.B) {
	for _, c := range testcases {
		b.Run("testcase"+strconv.Itoa(c.id), func(b *testing.B) {
			got := IsPalindrome3(c.str)

			if c.expected != got {
				b.Fatalf("str: %s, expected: %v, got: %v", c.str, c.expected, got)
			}
		})
	}
}
