package main

import (
	"fmt"
	"os"
	"strings"
)

var colourMap = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
}

const reset = "\033[0m"

func main() {
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")

	var color, substring, targetText string

	// Handle Arguments
	switch len(os.Args) {
	case 1:
		fmt.Println("Usage: go run . [color] [substring] [text]")
		return
	case 2:
		targetText = os.Args[1]
	case 3:
		color = os.Args[1]
		targetText = os.Args[2]
		substring = targetText
	case 4:
		color = os.Args[1]
		substring = os.Args[2]
		targetText = os.Args[3]
	default:
		fmt.Println("Error: Too many arguments.")
		return
	}

	if strings.HasPrefix(color, "--color=") {
		color = strings.TrimPrefix(color, "--color=")
	}

	ansiColor, colorExists := colourMap[color]
	inputRows := strings.Split(targetText, "\\n")

	for _, word := range inputRows {
		if word == "" {
			fmt.Println()
			continue
		}

		// 1. Create a "Color Map" for the word
		// This slice will keep track of which character indices need color
		shouldColor := make([]bool, len(word))
		if substring != "" && colorExists {
			searchStr := word
			offset := 0
			for {
				idx := strings.Index(searchStr, substring)
				if idx == -1 {
					break
				}
				// Mark all characters in this instance as 'true'
				for k := 0; k < len(substring); k++ {
					shouldColor[offset+idx+k] = true
				}
				// Move past this match to find the next one
				newStart := idx + len(substring)
				searchStr = searchStr[newStart:]
				offset += newStart
			}
		}

		// 2. Print ASCII Rows
		for r := 1; r <= 8; r++ {
			rowOutput := ""
			for j := 0; j < len(word); j++ {
				charIdx := int(word[j]-32)*9 + r

				if charIdx >= 0 && charIdx < len(lines) {
					charLine := lines[charIdx]

					// Check our boolean map for index 'j'
					if shouldColor[j] {
						rowOutput += ansiColor + charLine + reset
					} else {
						rowOutput += charLine
					}
				}
			}
			fmt.Println(rowOutput)
		}
	}
}
