package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {
	text := ""
	// if len(str) == 0 || len(str) < 3 {
	// 	return "\n"
	// }
	for i := 0; i < len(str); i++ {
		if (i+1)%3 == 0 {
			text += string(str[i])
		}
	}
	return fmt.Sprintf("%s\n", text)
}

// func main() {
// 	fmt.Print(ThirdTimeIsACharm("123456789"))
// 	fmt.Print(ThirdTimeIsACharm(""))
// 	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
// 	fmt.Print(ThirdTimeIsACharm("12"))
// }
