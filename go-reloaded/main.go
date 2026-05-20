package main

import (
	"fmt"
	"os"

	"go-reloaded/text_editor"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <inputFile> <outputFile>")
		return
	}

	text := text_editor.ReadFile(os.Args[1])

	text = text_editor.ProcessBrackets(text)

	text = text_editor.RemoveBrackets(text)
	text = text_editor.TransformAToAn(text)
	text = text_editor.FixPunctuation(text)
	text = text_editor.FixQuotesAndSpaces(text)

	text_editor.WriteFile(os.Args[2], text)

	fmt.Println("Processing complete! Output saved to", os.Args[2])
}
