package main

import "fmt"

func getCharLines(lines []string, index int) []string {
	start := index*4 + 1
	
	end := start + 3
	if end > len(lines) {
		return []string{}
	}
	return lines[start:end]
}

func main() {
	lines := []string{"", "A1", "A2", "A3", "", "B1", "B2", "B3"}
	fmt.Println(getCharLines(lines, 0))
	fmt.Println(getCharLines(lines, 1))
	fmt.Println(getCharLines(lines, 5))
}
