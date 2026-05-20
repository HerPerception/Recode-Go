package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input file> <output file>")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file, check if file exists.")
		return
	}
	text := Base(string(data))
	text1 := Case(text)
	text2 := FixQuotes(text1)
	text3 := AToAn(text2)
	text4 := FixPunct(text3)

	err = os.WriteFile(os.Args[2], []byte(text4), 0644)
	if err != nil {
		fmt.Println("Error writing file.")
		return
	} else {
		fmt.Printf("File processed successfully to %s\n", os.Args[2])
	}
}
