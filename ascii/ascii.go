package fs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// bannerMap is a map that stores the ASCII art for different banner files
var bannerMap map[string]string

// init initializes the bannerMap and loads the ASCII art from the banner files
func init() {
	bannerMap = make(map[string]string)
	loadBanner("standard.txt")
	loadBanner("shadow.txt")
	loadBanner("thinkertoy.txt")
}

// loadBanner reads the contents of a banner file and stores it in the bannerMap
func loadBanner(filename string) {
	filePath := filepath.Join("./banner", filename)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		return
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	bannerMap[filename] = strings.Join(lines, "\n")
}

// GetLetterArray retrieves the ASCII art representation for a given character from the specified banner file
func GetLetterArray(char rune, bannerStyle string) []string {
	
	banner, ok := bannerMap[bannerStyle]
	if !ok {
		fmt.Println("File doesn't exist")
		os.Exit(0)
	}
	alphabet := strings.Split(banner, "\n")
	start := (int(char) - 32) * 9
	if start < 0 || start >= len(alphabet) {
		return []string{}
	}
	arr := alphabet[start : start+9]
	return arr
}

// PrintAscii prints the ASCII art representation of a given string
func PrintAscii(str, bannerStyle string) {
	lines := strings.Split(str, "\n")
	letters := [][]string{}
	for _, line := range lines {
		for _, letter := range line {
			if letter < 32 || letter > 126 {
				fmt.Printf("Non-ASCII character '%c' encountered\n", letter)
				return
			}
			arr := GetLetterArray(letter, bannerStyle)
			if len(arr) == 0 {
				fmt.Println("File doesn't exist")
				return
			}
			letters = append(letters, arr)
			if letter == '\n' {
				fmt.Println()
			}
		}
	}
	// Print the ASCII art vertically
	for i := 1; i < 9; i++ {
		for _, letter := range letters {
			if len(letter) < i {
				fmt.Println("Error: File content modified")
				return
			}
			fmt.Printf("%s", letter[i])
		}
		fmt.Println()
	}
}
