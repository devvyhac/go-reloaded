package utils

import (
	"log"
	"os"
)

func ReadFile(file string) []byte {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal()
		os.Exit(1)
	}

	return data
}
