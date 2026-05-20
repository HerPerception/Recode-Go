package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) string {
	if len(str1) == 0 && len(str2) == 0 {
		return ""
	}
	i := 0
	text := ""
	for str1 <= str2 {
		if str1[i] != str2[i] {
			text += string(str1[i]) + string(str2[i])
		}
	}
	return text
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
