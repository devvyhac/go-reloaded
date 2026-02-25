package utils

func ParseVowel(words []string, index int) ([]string, int) {
	if index == len(words)-1 {
		return words, index
	}

	nextWord := []rune(words[index+1])
	char := lower(nextWord[0])

	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		words[index] += "n"
	}
	return words, index
}

func lower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + 32
	}
	return char
}
