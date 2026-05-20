package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Base(s string) string {
	str := strings.Fields(s)
	num := ""
	num2 := ""
	if len(str) != 3 {
		return "Invalid command. Check and try again."
	}
	for i := range str {
		if str[i] == "bin" {
			temp, err := strconv.ParseInt(str[i-1], 2, 64)
			if err != nil {
				return fmt.Sprintf("%s is not valid binary.", str[i-1])
			} else {
				return fmt.Sprintf("Decimal: %d", temp)
			}
		} else if str[i] == "hex" {
			temp, err := strconv.ParseInt(str[i-1], 16, 64)
			if err != nil {
				return fmt.Sprintf("%s is not valid hex", str[i-1])

			} else {
				return fmt.Sprintf("Decimal: %d", (temp))
			}
		} else if str[i] == "dec" {
			temp, err := strconv.ParseInt(str[i-1], 10, 64)
			if err != nil {
				return fmt.Sprintf("%s is not valid decimal", str[i-1])
			} else {
				num = strconv.FormatInt(temp, 2)
				num2 = strconv.FormatInt(temp, 16)
				res := fmt.Sprintf("Binary: %s", num)
	            res2 := fmt.Sprintf("Hex: %s", strings.ToUpper(num2))

	            return res + "\n" + res2
			}
		}
	}
	return fmt.Sprintf("%s does not include a valid decimal, hex or binary. Check input and try again.\n", s)
}

func main() {
	fmt.Println("Welcome.")

	for {
		fmt.Println("Enter Command below: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if input == "quit" {
			return
		} else {
			fmt.Println(Base(input))
		}
	}
}
