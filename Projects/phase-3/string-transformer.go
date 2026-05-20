package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StringTransformer(s string) string {
	str := strings.Fields(s)
	switch str[0] {
	case "upper":
		return strings.ToUpper(strings.Join(str[1:], " "))

	case "lower":
		return strings.ToLower(strings.Join(str[1:], " "))

	case "cap":
		for i := 0; i < len(str); i++ {
			str[i] = (strings.ToUpper(string(str[i][0])) + strings.ToLower(str[i][1:]))
		}
		return strings.Join(str[1:], " ")

	case "title":
		for i := range str {
			str[i] = strings.ToLower(str[i])
		}
		for i := range str {
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
		return strings.Join(str[1:], " ")

	case "snake":
		for i := range str {
			str[i] = strings.ToLower(str[i])
		}
		word := strings.Join(str[1:], "_")
		text := ""
		for _, ch := range word {
			if ch == '_' {
				text += string(ch)
			} else if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
				text += string(ch)
			}
		}
		return text

	case "reverse":

		text := ""
		for i := 0; i < len(str); i++ {
			index := str[i]
			for in := len(index) - 1; in >= 0; in-- {
				text += string(index[in])
			}
			str[i] = text
			text = ""
		}
		return strings.Join(str[1:], " ")
	}

	return fmt.Sprintf("Unknown command: %s, Valid Commands: upper, lower, cap, title, snake, reverse, exit", str[0])
}

func main() {
	fmt.Println("Welcome To Sentinel String Transformer.")
	for {
		fmt.Println("Enter Prompt below: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye...")
			return
		} else {
			if len(input) == 0 {
				continue
			}
			input = strings.ToLower(input)
			newInput := strings.Fields(input)
			if len(newInput) < 2 {
				fmt.Println("No text provided. Usage: <command> <text>")
			}
			fmt.Println(StringTransformer(input))
		}

	}
}
