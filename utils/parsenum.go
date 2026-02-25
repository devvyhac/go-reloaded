package utils

import (
	"log"
	"os"
	"slices"
	"strconv"
)

func ParseNum(words []string, index int, base int) ([]string, int) {
	num, err := strconv.ParseInt(words[index-1], base, 32)

	if err != nil {
		log.Fatal()
		os.Exit(1)
	}

	words[index-1] = strconv.Itoa(int(num))
	words = slices.Delete(words, index, index+1)
	index--
	return words, index
}
