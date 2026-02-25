package utils

import "slices"

func HandleCase(words []string, index int, action func(string) string) ([]string, int) {
	words[index-1] = action(words[index-1])
	words = slices.Delete(words, index, index+1)
	index--
	return words, index
}
