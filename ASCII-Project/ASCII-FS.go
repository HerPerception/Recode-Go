package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	length := len(os.Args)

	if length != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	} // This checks if the arguments includes the banner file.

	banners := ""

	if os.Args[2] == "standard" {
		banners = "standard.txt"
	} else if os.Args[2] == "shadow" {
		banners = "shadow.txt"
	} else if os.Args[2] == "thinkertoy" {
		banners = "thinkertoy.txt"
	}

	str := os.Args[1]
	if len(str) == 0 {
		return
	} else if str == "\n" {
		fmt.Println()
		return
	} // Prints a new line and returns if the string contains ONLY (\n)

	data, err := os.ReadFile(banners)
	if err != nil {
		return
	}
	slice := strings.Split(str, `\n`) /* This splits the string wherever it see (\n), not a newline. The backticks tells the
	compiler that this is a string literal, not a control character that represents newline in Go. */

	lines := strings.Split(string(data), "\n") /* This splits at every newline, the opposite of the first. */

	for i := range lines {
		lines[i] = strings.ReplaceAll(lines[i], "\r", "")
	} /* The thinkertoy.txt file use the Windows standard for ending lines so after splitting the files by new lines, it is
	necessary to remove the \r or it would cause characters to be overwritten when you print with thinkertoy.txt. */
	for c := range slice {
		if slice[c] == "" {
			fmt.Println()
			continue
		} /* Print a newline and continue if thereis an empty string in the text slice. */

		for r := 1; r < 9; r++ {
			row := []string{}

			s := slice[c]
			for i := 0; i < len(s); i++ {
				index := int(s[i]-32)*9 + r // Get the position of the character in the banner file.
				if index >= 0 && index < len(lines) {
					row = append(row, lines[index])
				} /* Append the characters of the string row by row, making sure the position is within the range of the
				slice of characters - lines. */
			}
			fmt.Println(strings.Join(row, "")) // Prints the characters row after row
		}
	}
}
