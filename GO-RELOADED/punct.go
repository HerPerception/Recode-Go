package main

import "strings"

func FixPunct(s string) string {
	str := strings.Fields(s)
	strs := strings.Join(str, " ")
	text := ""
	for i := 0; i < len(strs); i++ {
		if strs[i] == ' ' && strings.ContainsAny(string(strs[i+1]), "!.,;:?") {
			continue
		} else if i > 0 && strings.ContainsAny(string(strs[i-1]), "!.,;:?") && strings.ContainsAny(string(strs[i]), "!.,;:?") {
			text += string(strs[i])
		} else if i > 0 && strings.ContainsAny(string(strs[i-1]), "!.,;:?") && strs[i] != ' ' {
			text += " " + string(strs[i])
		} else {
			text += string(strs[i])
		}
	}
	return text
}
