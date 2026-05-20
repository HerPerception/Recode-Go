package main

import (
	"fmt"
	"strconv"
)

func lengthConvert(s string) string {
	start := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'm' && s[i-1] == 'k' {
			start = i - 1
			break
		} else if s[i] == 'm' && s[i-1] != 'k' {
			start = i
			break
		}
	}
	conv := s[start:]
	newValue := ""
	num, err := strconv.Atoi(s[:start])
	if err != nil {
		fmt.Println("Error with conversion. Check Input and try again")
	} else {
		if conv == "m" {
			n := float64(num) / float64(1000)
			newValue = strconv.FormatFloat(n, 'f', 3, 64) + "km"
		} else if conv == "km" {
			m := num * 1000
			newValue = strconv.Itoa(m) + "m"

		}
	}
	return newValue
}
func main() {
	fmt.Println(lengthConvert("105km"))
}
