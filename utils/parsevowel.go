package utils

import "unicode"

func ParseVowel(words []string, index int) ([]string, int) {
	if index == len(words)-1 {
		return words, index
	}

	nextWord := []rune(words[index+1])
	char := unicode.ToLower(nextWord[0])

	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		words[index] += "n"
	}
	return words, index
}
