package process

import (
	"fmt"
	"strings"
)

func AsciiArt(input string, banner [][]string) {

	lines := strings.Split(input, "\\n")

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char := range line {
				index := int(char) - 32

				if index < 0 || index >= len(banner) {
					index = 31 // replace with question mark character if character is not printable
				}

				fmt.Print(banner[index][row])
			}

			fmt.Println()
		}
	}
}
