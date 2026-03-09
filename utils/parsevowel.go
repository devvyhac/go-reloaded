package utils

import (
	"strings"
	"unicode"
)

func ParseVowel(words []string, index int) ([]string, int) {
	n := 1
	if index+n < len(words)-1 {
		for strings.HasPrefix(words[index+n], "(") || strings.HasSuffix(words[index+n], ")") {
			n++
		}
	}

	nextWord := []rune(words[index+n])
	char := unicode.ToLower(nextWord[0])

	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		words[index] += "n"
	}

	return words, index
}
