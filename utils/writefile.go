package utils

import (
	"log"
	"os"
)

func WriteFile(filename string, data []byte) {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal()
		os.Exit(1)
	}
}
