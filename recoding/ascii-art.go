package main

import (
	"fmt"
	"os"
	"strings"
)

func getCharLines(lines []string, index int) []string {
	start := index*9
	end := start + 8

	if start >= len(lines) || end > len(lines) {
		return []string{"", "", ""} 
	}

	return lines[start:end]
}

func renderWord(lines []string, word string) {
	for row := 0; row < 8; row++ {
		line := ""

		for _, char := range word {
			index := int(char) - ' '
			getLines := getCharLines(lines, index)
			line += getLines[row]
		}
		fmt.Println(line)
	}
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	renderWord(lines, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")


	fmt.Println(getCharLines(lines, 0))
}

//t.Fatalf("unexpected error loading standard.txt: %v", err)