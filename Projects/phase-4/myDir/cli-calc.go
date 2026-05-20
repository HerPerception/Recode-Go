package myDir

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculator(s string) string {
	str := strings.Fields(s)
	//var output string
	switch str[0] {
	case "add":
		Input = s
		first, err := strconv.Atoi(str[1])
		if err != nil {
			return "Error in first number, check input and try again"
		}
		sec, err := strconv.Atoi(str[2])
		if err != nil {
			return "Error in second number, check input and try again"
		}
		Output = fmt.Sprintf("result: %d", first+sec)
		return Output

	case "sub":
		Input = s
		first, err := strconv.Atoi(str[1])
		if err != nil {
			return "Error in first number, check input and try again"
		}
		sec, err := strconv.Atoi(str[2])
		if err != nil {
			return "Error in second number, check input and try again"
		}
		Output = fmt.Sprintf("result: %d", first-sec)
		return Output

	case "mul":
		Input = s
		first, err := strconv.Atoi(str[1])
		if err != nil {
			return "Error in first number, check input and try again"
		}
		sec, err := strconv.Atoi(str[2])
		if err != nil {
			return "Error in second number, check input and try again"
		}
		Output = fmt.Sprintf("result: %d", first*sec)
		return Output

	case "mod":
		Input = s
		first, err := strconv.Atoi(str[1])
		if err != nil {
			return "Error in first number, check input and try again"
		}
		sec, err := strconv.Atoi(str[2])
		if err != nil {
			return "Error in second number, check input and try again"
		}
		Output = fmt.Sprintf("result: %d", first%sec)
		return Output

	case "pow":
		Input = s
		first, err := strconv.Atoi(str[1])
		if err != nil {
			return "Error in first number, check input and try again"
		}
		sec, err := strconv.Atoi(str[2])
		if err != nil {
			return "Error in second number, check input and try again"
		}
		result := 1
		for i := 0; i < sec; i++ {
			result *= first
		}
		Output = fmt.Sprintf("result: %d", result)
		return Output

	case "last":
		if Output == "" {
			return "No previous session. Enter command <text> to begin one.\n"
		}
		if str[0] == "last" && Output != "" {
			Last = fmt.Sprintf("%s => %s", Input, Output)
		}

		return Last
	}
	//Last = fmt.Sprintf("%s => %s", Input, Output)

	if Last != "" {
		History = append(History, Last)
	}
	return "Invalid command. Check input and try again."
}
