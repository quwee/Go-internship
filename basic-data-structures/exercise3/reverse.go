package exercise3

import (
	"slices"
	"strings"
)

func ReverseStringWithoutLibs(s string) string {
	stringRunes := []rune(s)
	if len(stringRunes) == 0 && len(stringRunes) == 1 {
		return s
	}
	word := make([]rune, 0, 1)
	wordId := 0
	isPrevSpace := false

	for i := 0; i < len(stringRunes); i++ {
		if isSpace(stringRunes[i]) {
			if !isPrevSpace && i != 0 {
				word = reverseWord(word)

				for j := 0; j < len(word); j++ {
					stringRunes[wordId] = word[j]
					wordId++
				}
				wordId++
				word = make([]rune, 0, 1)
			} else {
				wordId++
			}
			isPrevSpace = true
		} else {
			isPrevSpace = false
			word = append(word, stringRunes[i])
		}
	}

	if !isPrevSpace {
		word = reverseWord(word)

		for j := 0; j < len(word); j++ {
			stringRunes[wordId] = word[j]
			wordId++
		}
	}

	return string(stringRunes)
}

func ReverseStringWithLibs(s string) string {
	words := strings.Fields(s)
	reversedWords := make([]string, len(words))

	for i, word := range words {
		wordChars := []rune(word)
		slices.Reverse(wordChars)
		reversedWords[i] = string(wordChars)
	}
	return strings.Join(reversedWords, " ")
}

func isSpace(r rune) bool {
	switch r {
	case ' ', '\t', '\n', '\r':
		return true
	}
	return false
}

func reverseWord(word []rune) []rune {
	result := make([]rune, len(word))

	for i := 0; i < len(word); i++ {
		result[i] = word[len(word)-1-i]
	}
	return result
}
