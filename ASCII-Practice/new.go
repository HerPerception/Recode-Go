package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	length := len(os.Args)

	if length < 2 || length > 3 {
		fmt.Println("Invalid number of Args")
		return
	}

	banners := "standard.txt"

	if length == 2 {
		banners = "standard.txt"
	}

	if length == 3 {
		banners = os.Args[2]
	}
	// if len(os.Args) != 3 {
	// 	fmt.Println("Include txt file.")
	// 	return
	// }
	str := os.Args[1]
	data, err := os.ReadFile(banners)
	if err != nil {
		return
	}
	slice := strings.Split(str, "\\n")
	lines := strings.Split(string(data), "\n")
	for i := range lines {
		lines[i] = strings.ReplaceAll(lines[i], "\r", "")

	}
	for c := range slice {
		for r := 1; r < 8; r++ {
			row := []string{}

			s := slice[c]
			for i := 0; i < len(s); i++ {
				index := int(s[i]-32)*9 + r
				if index >= 0 && index < len(lines) {
					row = append(row, lines[index])
				}
			}
			fmt.Println(strings.Join(row, ""))
		}
	}
}
