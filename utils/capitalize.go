package utils

func Capitalize(word string) string {
	if word == "" {
		return ""
	}

	runes := []rune(word)
	runes[0] -= 32
	return string(runes)
}
