package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide file name.")
		return
	}
	str, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file. Please make sure file exists")
		return
	} else {
		fmt.Println(string(str))
	}
}
