package utils

func ParseVowel(words []string, index int) ([]string, int) {
	nextWord := ""
	if len(words[index:]) > 0 {
		nextWord = words[index+1]
	}

	char := lower([]rune(nextWord)[0])

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
