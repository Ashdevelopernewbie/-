package bytes

import (
	"log"
	"os"
)

func CheckBytes(file string) int {
	output, err := os.Stat(file)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	return int(output.Size())
}

func Try(word string) string {
	return word
}
