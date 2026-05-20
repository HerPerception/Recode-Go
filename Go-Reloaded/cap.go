package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide input and output files.")
		return
	}
	input := os.Args[1]

	raw, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file.", err)
		return
	}

	text := string(raw)
	words := strings.Fields(text)

	//fmt.Println(words)

	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "(cap)" {
			if len(result) > 0 {
				last := result[len(result)-1]
				result[len(result)-1] = strings.Title(last)
			}
			continue
		}
		result = append(result, word)
	}
	final := strings.Join(result, " ")
	fmt.Println(final)
}
