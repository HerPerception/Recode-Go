package main

import (
	"fmt"
	"os"
)

func main(){
	first, second, third := os.Args[1], os.Args[2], os.Args[3]
	if len(os.Args) != 4 {
		fmt.Println()
		return
	}
	if len(second) != 1 || len(third) != 1 {
		fmt.Println(first)
		return
	}
	word := ""
	for _, ch := range first {
		if string(ch) == second {
			word += third
		} else {
			word += string(ch)
		}
	}
	fmt.Printf("%q\n", word)
}