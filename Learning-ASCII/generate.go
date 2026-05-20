package main

import "strings"

func Generate(text string, banner map[rune][]string) string {
	if len(text) == 0 {
		return ""
	}
	ch, err := Validate(text)
	if ch != 0 && err != nil {
		return "input has unsupported characters"
	}
	userText := SplitInput(text)
	newText := ""
	allEmpty := true
	for _, ch := range userText {
		if ch != "" {
			allEmpty = false
			break
		}
	}
	if allEmpty {
		for i := 0; i < len(userText)-1; i++ {
			newText += "\n"
		}
	} else {
		for _, ch := range userText {
			if ch == "" {
				newText += "\n"
				continue
			}
			newText += strings.Join(RenderLines(ch, banner), "\n")
		}
	}
	return newText
}
