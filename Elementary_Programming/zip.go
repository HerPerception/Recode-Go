package main

import (
	"fmt"
)

func ZipString(str string) string {
	text := ""
	count := 1
	for i := range str {
		if i < len(str)-1 && str[i] == str[i+1] {
			//next := i
			count++
		}
		if i < len(str)-1 && str[i] != str[i+1] {
			text += fmt.Sprintf("%s,%d", string(str[i]), count)
			//nexts := i + count
		}
		i++

	}
	return text
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
