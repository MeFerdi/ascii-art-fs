package main

import (
	"fmt"
	"os"
	"strings"

	fs "fs/ascii" // Import the ascii package from the local directory
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run main.go <input_string> [banner_style]\n\nEX: go run . something standard")
		return
	}
	str := os.Args[1]
	bannerStyle := "../banner/standard.txt" // Default banner style
	if len(os.Args) == 3 {
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
	if str == "" {
		return
	}

	modifiedStr, hasError := fs.SpecialCharacters(str)
	if hasError {
		fmt.Println(modifiedStr)
		return
	}

	str = modifiedStr
	str = strings.ReplaceAll(str, "\n", "\\n")
	lines := strings.Split(str, "\\n")

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
