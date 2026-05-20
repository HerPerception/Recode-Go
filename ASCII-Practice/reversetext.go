package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	// 1. Get input file and clean the banner file
	filename := strings.TrimPrefix(os.Args[1], "--reverse=")
	inputData, _ := os.ReadFile(filename)
	bannerData, _ := os.ReadFile("standard.txt")

	// Split input into 8 rows
	inputRows := strings.Split(strings.ReplaceAll(string(inputData), "\r", ""), "\n")
	if len(inputRows) < 8 {
		return
	}

	// Split banner into individual lines
	bannerLines := strings.Split(strings.ReplaceAll(string(bannerData), "\r", ""), "\n")

	result := ""
	cursor := 0
	maxLen := len(inputRows[0])

	// 2. The Sliding Window: Move horizontally across the columns
	for cursor < maxLen {
		found := false

		// Try different character widths (from 1 to 15 characters wide)
		for width := 1; width <= 15 && cursor+width <= maxLen; width++ {

			// Build the 8-line "block" from our input at the current position
			var currentBlock []string
			for r := 0; r < 8; r++ {
				currentBlock = append(currentBlock, inputRows[r][cursor:cursor+width])
			}

			// 3. Compare this block to every character in standard.txt
			// Each character in standard.txt starts at (Index * 9) + 1
			for charCode := 32; charCode <= 126; charCode++ {
				bannerStart := (charCode-32)*9 + 1
				match := true

				for r := 0; r < 8; r++ {
					if currentBlock[r] != bannerLines[bannerStart+r] {
						match = false
						break
					}
				}

				// 4. If it matches, we found the character!
				if match {
					result += string(rune(charCode))
					cursor += width // Move the cursor forward
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		// If no character matches, just move the cursor to avoid infinite loop
		if !found {
			cursor++
		}
	}

	fmt.Println(result)
}
