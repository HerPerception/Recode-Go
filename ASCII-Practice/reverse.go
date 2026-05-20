package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	flag := os.Args[1]
	file := strings.TrimPrefix(flag, "--reverse=")
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error")
		return
	}
	str := strings.Split(string(data), "\n")
	//fmt.Println(str)
	// data2, err := os.ReadFile("standard.txt")
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// }
	// in := 0
	// text := ""
	// fmt.Println(len(str))
	// lines := strings.Split(string(data2), "\n")

	for j := range str {
		// start := j * 8
		// end := start + 7
		// for i := range lines {
		// 	if j+8 < len(str) && strings.Join(str[j:j+7], "") == lines[i] {
		// 		in = (i / 9) + 32
		// 		text += string(rune(in))
		// 		//fmt.Println(lines[i])
		// 	}
		// }

		// in = 0
		fmt.Println(str[j])
	}
	//fmt.Println(text)
}
