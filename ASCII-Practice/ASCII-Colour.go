package main

import (
	"fmt"
	"os"
	"strings"
)

var colourMap = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Error")
		return
	}
	str := ""
	if len(os.Args) == 2 {
		str = os.Args[1]
	} else if len(os.Args) == 3 {
		str = os.Args[2]
	} else if len(os.Args) == 4 {
		str = os.Args[3]
	}
	substr := str
	if len(os.Args) == 4 {
		substr = os.Args[2]

	}
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file.")
		return
	}
	var indexes []string
	text := ""
	lines := strings.Split(string(data), "\n")
	actualColour := strings.Split(os.Args[1], "=")[1]
	slice := strings.Split(str, " ")
	for i := 0; i < len(slice); i++ {
		s := slice[i]
		for j := 0; j < len(s); j++ {
			if j < j+len(substr) && len(substr) <= len(s) && s[j] == substr[0] && s[j+(len(substr)-1)] == substr[len(substr)-1] {
				_, ok := colourMap[actualColour]
				if ok {
					in := j
					for in < len(substr) {
						text += colourMap[actualColour] + string(s[in]) + "\033[0m"
						in++
					}
				}
				j += len(substr) - 1
			} else {
				text += string(s[j])
			}
		}
		indexes = append(indexes, text)
		text = ""
	}
	indexes = fmt.Sprint(indexes)
	for i := 0; i < len(indexes); i++ {
		s := indexes[i]
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
