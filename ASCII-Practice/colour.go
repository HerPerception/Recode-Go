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
		fmt.Println("Invalid number of arguments.")
		return
	}
	// data, err := os.ReadFile("standard.txt")
	// if err != nil {
	// 	fmt.Println("Error reading file.")
	// 	return
	// }
	// lines := strings.Split(string(data), "\n")
	color := os.Args[1]
	str := ""
	if len(os.Args) == 2 {
		str = os.Args[1]
	}
	substr := ""
	text := ""
	colour := strings.Split(color, "=")
	actualColour := colour[1]
	if len(os.Args) == 3 {
		str = os.Args[2]
		substr = str
		text += colourMap[actualColour] + substr + "\033[0m"
		fmt.Println(text)
		return
	}
	if len(os.Args) == 4 {
		str = os.Args[3]
		substr = os.Args[2]
		//fmt.Println(actualColour, substr, str)
		slice := strings.Split(str, " ")
		var indexes []string
		// start := 0
		// end := 0
		//text := ""
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
		fmt.Println(strings.Join(indexes, " "))
	}
}
