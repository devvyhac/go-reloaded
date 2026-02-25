package utils

import "slices"

func ParsePunctuation(words []string, index int) ([]string, int) {
	puncs := map[rune]bool{'.': true, ',': true, '?': true, '!': true, ':': true, ';': true}

	runes := []rune(words[index])

	if len(runes) == 0 || !puncs[runes[0]] || index == 0 {
		return words, index
	}

	if len(runes) == 1 {
		words[index-1] += string(runes[0])
		words = slices.Delete(words, index, index+1)
		index--
		return words, index
	}

	for i, char := range runes {
		// ,!?
		if !puncs[char] {
			words[index] = string(runes[i:])
			words[index-1] += string(runes[:i])
			return words, index
		}
	}

	words[index-1] += words[index]
	words = slices.Delete(words, index, index+1)
	index--
	return words, index
}
