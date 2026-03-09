package main

import (
	"fmt"
	"os"
	"reloaded/utils"
)

func main() {
	sample, result := os.Args[1], os.Args[2]
	fmt.Println(sample, result)

	input := utils.ReadFile(sample)

	sentence := string(input)
	fmt.Println(sentence)
	output := utils.ParseInput(sentence)

	utils.WriteFile(result, output)
}
