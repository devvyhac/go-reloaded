package main

import (
	"os"
	"reloaded/utils"
)

func main() {
	sample, result := os.Args[1], os.Args[2]

	input := utils.ReadFile(sample)

	sentence := string(input)
	output := utils.ParseInput(sentence)

	utils.WriteFile(result, output)
}
