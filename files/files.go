package files

import (
	"log"
	"os"
)

func WriteFile(path string, data []byte) {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
