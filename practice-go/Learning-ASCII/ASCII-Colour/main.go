package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}
	text := ""
	substr := ""
	banner := "standard.txt"
	if len(os.Args) == 3 {
		text = os.Args[2]
		substr = os.Args[2]
	}
	if len(os.Args) == 4 {
		text = os.Args[3]
		substr = os.Args[2]
	}
	if len(os.Args) == 5 {
		text = os.Args[3]
		substr = os.Args[2]
		banner = os.Args[4]
	}
	result, err := LoadBanner(banner)
	if err != nil {
		fmt.Println(err)
		return
	}
		//substr := os.Args[2]
	fmt.Println(Generate(text, result, substr))

}
