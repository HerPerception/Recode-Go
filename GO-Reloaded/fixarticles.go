package main

import "strings"

func AToAn(s string) string {
	str := strings.Fields(s)
	for i := 0; i < len(str)-1; i++ {
		if str[i] == "a" && strings.ContainsAny(string(str[i+1][0]), "aeiouhAEIOUH") {
			str[i] = "an"
		} else if str[i] == "A" && strings.ContainsAny(string(str[i+1][0]), "aeiouhAEIOUH") {
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
		if s[i] == '\'' && inQuote == false {
			text += string(s[i])
			inQuote = true
		} else if inQuote && s[i] == ' ' && s[i-1] == '\'' {
			continue
		} else if inQuote == true && s[i] == ' ' && s[i+1] == '\'' {
			continue
		} else if inQuote && strings.ContainsAny(string(s[i-1]), ".,!;'") && s[i] == ' ' && s[i+1] == '\'' {
			continue
		} else if inQuote && !strings.ContainsAny(string(s[i-1]), ".,!;'") && s[i] == '\'' && s[i-1] != ' ' {
			text += string(s[i])
		} else {
			text += string(s[i])
		}
	}
	return text
}
