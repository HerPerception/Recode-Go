package main

import (
	"errors"
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}
	_, err := ValidateInput(input)
	if err != nil {
		return "", errors.New("unsupported character in input text")
	}
	slice := SplitInput(input)
	word := []string{}

	allEmpty := true
	for _, ch := range slice {
		if ch != "" {
			allEmpty = false
		}
		if ch == "" {
			word = append(word, "")
			continue
		}
		word = append(word, strings.Join(RenderLine(ch, banner), "\n"))
	}
	if allEmpty {
		stri := ""
		for i := 0; i < len(slice)-1; i++ {
			stri += "\n"
		}
		return stri, nil
	}
	return strings.Join(word, "\n") + "\n", nil
}
