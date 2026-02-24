package main

import (
	"fmt"
	"log"
	"os"
	"reloaded/utils"
	"strconv"
	"strings"
)

func main() {
	sample := os.Args[1]

	input, err := os.ReadFile(sample)
	if err != nil {
		log.Fatal()
		os.Exit(1)
	}

	sentence := string(input)
	ParseInput(sentence)

	// errs := os.WriteFile(result, input, 0644)
	// if errs != nil {
	// 	log.Fatal()
	// }

}

func ParseInput(data string) {
	words := strings.Fields(data)

	for i, word := range words {
		switch word {

		case "(cap,":
			ParseMultiple(words, i, "cap")

		case "(cap)":
			words[i-1] = utils.Capitalize(words[i-1])
			words[i] = ""
			// words = slices.Delete(words, i, i+1)

		case "(up,":
			ParseMultiple(words, i, "up")

		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words[i] = ""

		case "(low,":
			ParseMultiple(words, i, "low")

		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words[i] = ""
		}

	}

	fmt.Println(words)
}

func ParseMultiple(words []string, i int, mode string) {
	num, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
	if err != nil {
		log.Fatal()
		os.Exit(1)
	}

	for num > 0 {
		fmt.Println(num)
		switch mode {
		case "low":
			words[i-num] = strings.ToLower(words[i-num])
		case "up":
			words[i-num] = strings.ToUpper(words[i-num])
		case "cap":
			words[i-num] = utils.Capitalize(words[i-num])
		}

		num--
	}
	words[i] = ""
	words[i+1] = ""
}
