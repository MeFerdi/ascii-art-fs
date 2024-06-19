package main

import (
	"fmt"
	"os"
	"strings"

	fs "fs/ascii" // Import the ascii package from the local directory
)

func main() {
	// Check if the command-line arguments are valid
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run main.go <input_string> [banner_style]\n\nEX: go run . something standard")
		return
	}

	// Get the input string and banner style from the command-line arguments
	str := os.Args[1]
	bannerStyle := "standard.txt" // Default banner style
	if len(os.Args) == 3 {
		// Check if the banner style is a valid file
		if !strings.HasSuffix(os.Args[2], ".txt") {
			bannerStyle = os.Args[2] + ".txt"
		} else {
			bannerStyle = os.Args[2]
			if !fs.IsValidBanner(bannerStyle) {
				fmt.Println("Error: Banner style invalid")
				return
			}
		}
	}

	// Check if the input string is empty
	if str == "" {
		return
	}

	// Handle special characters in the input string
	modifiedStr, hasError := fs.SpecialCharacters(str)
	if hasError {
		fmt.Println(modifiedStr)
		return
	}

	// Replace newline characters with "\\n" and split the string into lines
	str = modifiedStr
	str = strings.ReplaceAll(str, "\n", "\\n")
	lines := strings.Split(str, "\\n")

	// Print the ASCII art for each line
	newlineCount := 0
	for _, line := range lines {
		if line == "" {
			newlineCount++
			if newlineCount > 0 {
				fmt.Println()
			}
		} else {
			fs.PrintAscii(line, bannerStyle) // Call the PrintAscii function from the ascii package to print ASCII art
			fmt.Println()
		}
	}
}
