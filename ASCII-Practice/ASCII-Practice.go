package main

import (
	"fmt"
	"os"
	"strings"
)
func main(){
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Invalid number of arguments.")
	}

	file := "standard.txt"
	if len(os.Args) == 2 {
		file = "standard.txt"
	}
	
	if len(os.Args) == 3 {
		file = os.Args[2]
	}
	if len(os.Args[1]) == 0 {
		return
	} else if os.Args[1] == `\n` {
		fmt.Println()
		return
	}

	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file.")
	}
	slice := strings.Split(os.Args[1], `\n`)
	lines := strings.Split(string(input), "\n")

	for i := range lines {
		lines[i] = strings.ReplaceAll(lines[i], "\r", "")
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] == `\n` {
			fmt.Println()
			continue
		}
		s := slice[i]
		for r := 1; r < 9; r++ {
			row := []string{}
			for j := 0; j < len(s); j++ {
				in := int(s[j] - 32) * 9 + r
				if in >= 0 && in < len(lines) {
					row = append(row, lines[in])
				}
			}
			fmt.Println(strings.Join(row, ""))
		}
	}
	
}