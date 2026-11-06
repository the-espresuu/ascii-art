package process

import (
	"os"
	"strings"
)

func FindStyle(input string) []byte {

	// Making input directory-readable
	input = strings.ToLower(input)
	if !strings.HasSuffix(input, ".txt") {
		input += ".txt"
	}
	input = "banners/" + input

	// Reading the file
	data, err := os.ReadFile(input)
	if err != nil {
		return nil
	}
	return data
}
