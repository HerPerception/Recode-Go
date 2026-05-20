package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide input file.")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file. Check if file exists.")
		return
	} else {
		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			fmt.Println(i, line)
		}

	}
}
