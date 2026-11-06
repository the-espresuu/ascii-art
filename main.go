package main

import (
	process "ascii/functions"
	"fmt"
	"os"
)

func main() {
	// Check for valid number of arguments
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}

	// Default values and input values
	bannerinput := "standard"
	input := os.Args[1]

	// Change the bannerinput if an input is found
	if len(os.Args) == 3 {
		bannerinput = os.Args[2]
	}

	// Main process and error handling
	style := process.FindStyle(bannerinput)
	if style == nil {
		fmt.Printf("Error: Font style not found: %s\n", bannerinput)
		return
	}
	banner := process.ParseCharacters(string(style))
	process.AsciiArt(input, banner)
}
