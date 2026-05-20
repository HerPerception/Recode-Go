package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . <text> <filename>")
		return
	}
	str := strings.Split(os.Args[1], "\\n")
	banner := "standard.txt"
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}
	data, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Error reading file.")
		return
	}
	if len(os.Args[1]) == 0 {
		return
	}
	if os.Args[1] == "\n" {
		fmt.Println()
		return
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	for i := 0; i < len(str); i++ {
		s := str[i]
		// if s == "" {
		// 	fmt.Println()
		// 	continue
		// }
		for r := 1; r < 9; r++ {
			row := []string{}
			for j := 0; j < len(s); j++ {
				index := int(s[j]-32)*9 + r
				if index > 0 && index < len(lines) {
					row = append(row, lines[index])
				}
			}
			fmt.Println(strings.Join(row, ""))
		}
	}
}
