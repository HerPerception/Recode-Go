package text_editor

import (
	"strconv"
	"strings"
)

// ProcessBrackets handles all (up), (low), (cap), (hex), (bin) commands
func ProcessBrackets(text string) string {
	starts := findIndexes(text, '(')
	ends := findIndexes(text, ')')

	// Make sure we have matching brackets
	if len(starts) != len(ends) {
		return text
	}

	for i := 0; i < len(starts); i++ {
		cmdStr := text[starts[i]+1 : ends[i]]
		cmd, num := parseCommand(cmdStr)

		switch cmd {
		case "up", "low", "cap":
			text = ApplyCase(text, cmd, num)
		case "hex":
			text = ConvertHex(text)
		case "bin":
			text = ConvertBinary(text)
		}
	}

	return text
}

// Returns indexes of a character in a string
func findIndexes(text string, char rune) []int {
	var indexes []int
	for i, r := range text {
		if r == char {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// parse command like "up" or "up,3"
func parseCommand(s string) (string, int) {
	if strings.Contains(s, ",") {
		parts := strings.Split(s, ",")
		num := 1
		if len(parts) > 1 {
			numStr := strings.TrimSpace(parts[1])
			if n, err := strconv.Atoi(numStr); err == nil {
				num = n
			}
		}
		return strings.TrimSpace(parts[0]), num
	}
	return strings.TrimSpace(s), 1
}
