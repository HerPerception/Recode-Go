package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(data)
	split := strings.Split(string(data), "\n")
	split = split[1:]
	// for i, ch := range split {
	// 	fmt.Println(i, ch)
	// }btd

	var myMap = make(map[rune][]string)
	for i := ' '; i <= '~'; i++ {
		start := int(i-32) * 9
		end := start + 8
		myMap[i] = split[start:end]
	}

	a := myMap['a']
	for i := range a {
		fmt.Println(a[i])
	}
}
