package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Too many arguments.")
		return
	}
	banner := ""
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}
	if len(os.Args) == 2 {
		banner = "standard.txt"
	}

	data, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Error reading file.")
		return
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	str := strings.Split(os.Args[1], "\\n")

	var newText string
	for i := range str {
		s := str[i]
		var text string
		for r := 1; r < 9; r++ {
			var row []string
			if s == "" {
				text = fmt.Sprintln()
				continue
			}
			for _, ch := range s {
				index := int(ch-32)*9 + r
				if index > 0 && index < len(lines) {
					row = append(row, lines[index])
				} else {
					continue
				}
			}
			text += fmt.Sprintln(strings.Join(row, ""))
		}
		newText += fmt.Sprint(text)

	}
	err = os.WriteFile("file.txt", []byte(newText), 0644)
	if err != nil {
		fmt.Println("Error writing file.")
	}
	//fmt.Print(newText)

}
