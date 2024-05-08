package exercise3

import (
	"strconv"
	"testing"
)

type testcase struct {
	id       int
	str      string
	expected string
}

var testcases = []testcase{
	{1, "abc abc abc", "cba cba cba"},
	{2, "abc cba abc cba", "cba abc cba abc"},
	{3, "123 abc", "321 cba"},
	{4, "123 abc lp", "321 cba pl"},
	{5, "ac wq sp st", "ca qw ps ts"},
	{6, "你好 世界 好你", "好你 界世 你好"},
	{7, "Café über", "éfaC rebü"},
}

func BenchmarkReverseStringWithoutLibs(b *testing.B) {
	for _, c := range testcases {
		b.Run("testcase_"+strconv.Itoa(c.id), func(b *testing.B) {
			got := ReverseStringWithoutLibs(c.str)

			if c.expected != got {
				b.Fatalf("str: %s, expected: %s, got: %s", c.str, c.expected, got)
			}
		})
	}
}

func BenchmarkReverseStringWithLibs(b *testing.B) {
	for _, c := range testcases {
		b.Run("testcase_"+strconv.Itoa(c.id), func(b *testing.B) {
			got := ReverseStringWithLibs(c.str)

			if c.expected != got {
				b.Fatalf("str: %s, expected: %s, got: %s", c.str, c.expected, got)
			}
		})
	}
}
