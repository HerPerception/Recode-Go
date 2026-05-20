package text_editor

import (
	"fmt"
	"os"
)

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	return string(data)
}

func WriteFile(filename, text string) {
	err := os.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
