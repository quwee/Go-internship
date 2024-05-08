package exercise1

import (
	"regexp"
	"slices"
	"strings"
)

var reg = regexp.MustCompile("[^a-zA-Z0-9]+")

func IsPalindrome1(s string) bool {
	if len(s) == 0 {
		return false
	}
	s = reg.ReplaceAllString(strings.ToLower(s), "")
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func IsPalindrome2(s string) bool {
	if len(s) == 0 {
		return false
	}
	s = reg.ReplaceAllString(strings.ToLower(s), "")
	bytes := []byte(s)
	slices.Reverse(bytes)

	for i := 0; i < len(s); i++ {
		if s[i] != bytes[i] {
			return false
		}
	}
	return true
}

func IsPalindrome3(s string) bool {
	if len(s) == 0 {
		return false
	}
	s = reg.ReplaceAllString(strings.ToLower(s), "")
	return checkPalindromeRec(s)
}

func checkPalindromeRec(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return checkPalindromeRec(s[1 : len(s)-1])
}
