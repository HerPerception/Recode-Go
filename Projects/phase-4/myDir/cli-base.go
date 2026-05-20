package myDir

import (
	"fmt"
	"strconv"
	"strings"
)

func Base(s string) string {
	str := strings.Fields(s)
	switch str[0] {
	case "hex":
		Input = s
		num, err := strconv.ParseInt(str[1], 16, 64)
		if err != nil {
			return "Invalid hexadecimal. Check input and try again."
		} else {
			Output = fmt.Sprintf("Decimal: %d", num)
			return Output
		}

	case "bin":
		Input = s
		num, err := strconv.ParseInt(str[1], 2, 64)
		if err != nil {
			return "Invalid binary. Check input and try again."
		} else {
			Output = fmt.Sprintf("Decimal: %d", num)
			return Output
		}

	case "dec":
		Input = s
		num, err := strconv.ParseInt(str[1], 10, 64)
		if err != nil {
			return "Invalid decimal. Check input and try again."
		} else {
			num1 := strconv.FormatInt(num, 2)
			num2 := strconv.FormatInt(num, 16)
			res := fmt.Sprintf("Binary: %s", num1)
			res2 := fmt.Sprintf("Hexadecimal: %s", strings.ToUpper(num2))
			Output = res + "\n" + res2
			return Output
		}
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
	return fmt.Sprintf("%s is invalid Base Converter command. Check input and try again.", str[0])
}
