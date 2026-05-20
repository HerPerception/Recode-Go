package text_editor

import (
	"strings"
)

// ApplyCase modifies last 'num' words before the bracket command
func ApplyCase(text, mode string, num int) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "("+mode) {
			start := i - num
			if start < 0 {
				start = 0
			}
			for j := start; j < i; j++ {
				switch mode {
				case "up":
					words[j] = strings.ToUpper(words[j])
				case "low":
					words[j] = strings.ToLower(words[j])
				case "cap":
					words[j] = strings.Title(words[j])
				}
			}
		}
	}
	return strings.Join(words, " ")
}
