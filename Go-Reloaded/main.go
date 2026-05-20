package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide input and output files.")
		return
	}
	input := os.Args[1]
	//output := os.Args[2]
	raw, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file.", err)
		return
	}
	text := string(raw)
	sliced := strings.Fields(text)
	var result []string

	for i := 0; i < len(sliced); i++ {

		word := sliced[i]
		cleanWord := strings.TrimFunc(word, func(r rune) bool {
			return unicode.IsPunct(r)
		})
		switch cleanWord {

		case "(cap)":
			if i > 0 {
				result[len(result)-1] = strings.Title(result[len(result)-1])
			}
		case "up":
			if i > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
		case "low":
			if i > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}
		default:
			result = append(result, word)
		}
	}
	final := strings.Join(result, " ")

	fmt.Println(final)
}
