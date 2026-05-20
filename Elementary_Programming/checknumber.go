package main

import "fmt"

func checkNumber(arg string) bool {
	for _, ch := range arg {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkNumber("Hello"))
	fmt.Println(checkNumber("Hello1"))
	fmt.Println(checkNumber("H1e2l3l4o5"))
}