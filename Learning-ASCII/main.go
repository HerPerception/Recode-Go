package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2  || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println(`go run . --color=<color> <substring to be colored> "something"`)
		return
	}
	result, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var text string
	if len(os.Args) == 3 {
		text = os.Args[2]
	}
	if len(os.Args) == 4 {
		text = os.Args[3]
	}
	fmt.Println(Generate(text, result))
}
