package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func trimPunct(word string) string {
	return strings.TrimFunc(word, func(r rune) bool {
		return unicode.IsPunct(r)
	})
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide input and output files.")
		return
	}

	input := os.Args[1]
	raw, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(raw)
	sliced := strings.Fields(text)
	var result []string

	for _, word := range sliced {
		clean := trimPunct(word)

		switch clean {
		case "cap":
			if len(result) > 0 {
				result[len(result)-1] = strings.Title(result[len(result)-1])
			}
			continue // skip appending the control word
		case "up":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
			continue
		case "low":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}
			continue
		default:
			result = append(result, word) // keep the word as-is (with punctuation)
		}
	}

	final := strings.Join(result, " ")
	fmt.Println(final)
}
