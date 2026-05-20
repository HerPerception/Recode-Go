package main

import "strings"

func AToAn(s string) string {
	str := strings.Fields(s)
	for i := 0; i < len(str)-1; i++ {
		if str[i] == "a" && strings.ContainsAny(string(str[i+1][0]), "aeiouhAEIOH") {
			str[i] = "an"
		} else if str[i] == "A" && strings.ContainsAny(string(str[i+1][0]), "aeiouhAEIOH") {
			str[i] = "An"
		}
	}
	return strings.Join(str, " ")
}

func FixQuotes(s string) string {
	str := strings.Fields(s)
	s = strings.Join(str, " ")
	var text string
	inQuote := false
	for i := 0; i < len(s); i++ {
		if s[i] == '\'' {
			text += string(s[i])
			inQuote = true
		} else if s[i] == ' ' && s[i-1] == '\'' && inQuote {
			continue
		} else if inQuote && s[i] == ' ' && strings.ContainsAny(string(s[i-1]), "!.,;:?") && s[i+1] == '\'' {
			continue
		} else if inQuote && s[i] == ' ' && s[i+1] == '\'' && !strings.ContainsAny(string(s[i-1]), "!.,;:?") {
			continue
		} else if inQuote && s[i] == '\'' {
			text += string(s[i])
			inQuote = false
		} else {
			text += string(s[i])
		}
	}
	return text
}
