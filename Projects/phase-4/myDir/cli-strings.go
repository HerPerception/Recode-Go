package myDir

import (
	"fmt"
	"strings"
)

var Output string
var Input string
var History []string
var Last string

func StringsTran(s string) string {
	str := strings.Fields(s)
	switch str[0] {
	case "upper":
		Input = s
		for i := range str {
			str[i] = strings.ToUpper(str[i])
		}

		Output = strings.Join(str[1:], " ")
		return Output

	case "lower":
		Input = s
		for i := range str {
			str[i] = strings.ToLower(str[i])
		}
		Output = strings.Join(str[1:], " ")
		return Output

	case "cap":
		Input = s
		for i := 0; i < len(str); i++ {
			str[i] = (strings.ToUpper(string(str[i][0])) + strings.ToLower(str[i][1:]))
		}
		Output = strings.Join(str[1:], " ")
		return Output

	case "title":
		Input = s

		for i := range str {
			str[i] = strings.ToLower(str[i])
			if i > 1 && (str[i] == "a" || str[i] == "an" || str[i] == "the" || str[i] == "and") {
				continue
			} else if i > 1 && (str[i] == "but" || str[i] == "or" || str[i] == "for" || str[i] == "nor") {
				continue
			} else if i > 1 && (str[i] == "on" || str[i] == "at" || str[i] == "to" || str[i] == "by") {
				continue
			} else if i > 1 && (str[i] == "in" || str[i] == "of" || str[i] == "up" || str[i] == "as" || str[i] == "is" || str[i] == "it") {
				continue
			} else {
				str[i] = (strings.ToUpper(string(str[i][0])) + strings.ToLower(str[i][1:]))
			}
		}
		Output = strings.Join(str[1:], " ")
		return Output

	case "snake":
		Input = s
		strs := strings.ToLower(strings.Join(str[1:], "_"))

		word := ""
		for i := range strs {
			if strs[i] == '_' {
				word += string(strs[i])
			} else if strs[i] >= 'a' && strs[i] <= 'z' || strs[i] >= '0' && strs[i] <= '9' {
				word += string(strs[i])
			}
		}
		Output = word
		return Output

	case "reverse":
		Input = s
		text := ""
		for i := range str {
			index := str[i]
			for in := len(index) - 1; in >= 0; in-- {
				text += string(index[in])
			}
			str[i] = text
			text = ""
		}
		Output = strings.Join(str[1:], " ")
		return Output

	case "last":
		if Output == "" {
			return "No previous session. Enter command <text> to begin one.\n"
		}
		if str[0] == "last" && Output != "" {
			Last = fmt.Sprintf("%s => %s", Input, Output)
		}
		if Last != "" {
			History = append(History, Last)
		}
		return Last
	}
	// Last = fmt.Sprintf("%s => %s", Input, Output)
	History = append(History, Last)

	return fmt.Sprintf("%s is not a string transformer command. Check input and try agin.", str[0])
}

func Back(s string) {
	if s == "back" {
		return
	}
}
