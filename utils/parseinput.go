package utils

import (
	"fmt"
	"strings"
)

func ParseInput(data string) []byte {
	words := strings.Fields(data)
	start := -1
	end := -1

	for i := 0; i < len(words); i++ {
		switch words[i] {

		case "'":
			words, i, start, end = ParseQuote(words, i, start, end)

		case "'s", "'t", "'ve", "'re", "'ll":
			words, i = ParseApostrophe(words, i)

		case "(cap,":
			words, i = ParseMultiple(words, i, Capitalize)

		case "(hex)":
			words, i = ParseNum(words, i, 16)

		case "(bin)":
			words, i = ParseNum(words, i, 2)

		case "(cap)":
			words, i = HandleCase(words, i, Capitalize)

		case "(up,":
			words, i = ParseMultiple(words, i, strings.ToUpper)

		case "(up)":
			words, i = HandleCase(words, i, strings.ToUpper)

		case "(low,":
			words, i = ParseMultiple(words, i, strings.ToLower)

		case "(low)":
			words, i = HandleCase(words, i, strings.ToLower)

		case "a", "A":
			words, i = ParseVowel(words, i)

		default:
			words, i = ParsePunctuation(words, i)
		}

	}

	fmt.Printf("Lenght: %d, Capacity: %d\n", len(words), cap(words))
	result := strings.Join(words, " ")
	fmt.Println(result)
	return []byte(result)
}
