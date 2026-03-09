package utils

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ParseMultiple(words []string, i int, action func(string) string) ([]string, int) {

	num, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
	if err != nil {
		log.Fatal()
		os.Exit(1)
	}

	if i < num {
		num = i
	}

	for num > 0 {
		words[i-num] = action(words[i-num])
		num--
	}

	words = slices.Delete(words, i, i+2)
	i--
	return words, i
}
