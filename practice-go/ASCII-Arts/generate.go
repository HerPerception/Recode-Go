package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if len(input) == 0 {
		return ""
	}
	slice := SplitInput(input)
	word := []string{}

	allEmpty := false
	for _, ch := range slice {
		if ch != "" {
			allEmpty = true
		}
		if ch == "" {
			word = append(word, "")
			continue
		}
		word = append(word, strings.Join(RenderLine(ch, banner), "\n"))
	}
	if allEmpty == false {
		stri := ""
		for i := 0; i < len(slice)-1; i++ {
			stri += "\n"
		}
		return stri
	}
	return strings.Join(word, "\n") + "\n"
}
