package text_editor

import (
	"fmt"
	"strconv"
	"strings"
)

// Convert binary string before (bin) to decimal
func ConvertBinary(text string) string {
	words := strings.Fields(text)
	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" {
			if num, err := strconv.ParseInt(words[i-1], 2, 64); err == nil {
				words[i-1] = fmt.Sprintf("%d", num)
			}
		}
	}
	return strings.Join(words, " ")
}

// Convert hex string before (hex) to decimal
func ConvertHex(text string) string {
	words := strings.Fields(text)
	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" {
			if num, err := strconv.ParseInt(words[i-1], 16, 64); err == nil {
				words[i-1] = fmt.Sprintf("%d", num)
			}
		}
	}
	return strings.Join(words, " ")
}
