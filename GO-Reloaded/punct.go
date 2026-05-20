package main

import "strings"

func Punctuations(s string) string {
	str := strings.Fields(s)
	s = strings.Join(str, " ")
	var text string
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && strings.ContainsAny(string(s[i]), ".,!;:?") && strings.ContainsAny(string(s[i+1]), ".,!;:'") {
			text += string(s[i])
		} else if s[i] == ' ' && strings.ContainsAny(string(s[i+1]), ".,!;:?") {
			continue
		} else if i < len(s)-1 && strings.ContainsAny(string(s[i]), ".,!;:?") && s[i+1] != ' ' {
			text += string(s[i]) + string(' ')
		} else {
			text += string(s[i])
		}
	}
	return text
}
