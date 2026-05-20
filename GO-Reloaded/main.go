package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file.")
		return
	}
	text := Base(string(data))
	text = Case(text)
	text = AToAn(text)
	text = FixQuotes(text)
	text = Punctuations(text)

	err = os.WriteFile(os.Args[2], []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file.")
		return
	}

	fmt.Printf("File processed successfully. Saved to %s\n", os.Args[2])

}
