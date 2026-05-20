package main

import (
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

func RenderLines(text string, banner map[rune][]string) []string {
	actualColour := strings.TrimPrefix(os.Args[1], "--color=")
	const reset = "\033[0m"
	var slice []string
	var substr string
	if len(os.Args) == 3 {
		substr = text
	} else if len(os.Args) == 4 {
		substr = os.Args[2]
	}
	num := 0
	for r := 0; r < 8; r++ {
		var builder strings.Builder
		//j := 0
		for i, ch := range text {
			if  i < len(text)-1 && i +len(substr) <= len(text) && text[i] == substr[0] && text[i:i+len(substr)] == substr {
				num = len(substr)
			}
			if num > 0 {
				builder.WriteString(colourMap[actualColour] + banner[ch][r] + reset)
				num--
			} else {
				builder.WriteString(banner[ch][r])
			}
			//num = false
		}
		slice = append(slice, builder.String())
	}
	return slice
}
