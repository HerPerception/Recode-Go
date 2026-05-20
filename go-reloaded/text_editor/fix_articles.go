package text_editor

import (
	"strings"
	"unicode"
)

// Remove all brackets content
func RemoveBrackets(text string) string {
	var result strings.Builder
	inBracket := false
	for _, r := range text {
		switch r {
		case '(':
			inBracket = true
		case ')':
			inBracket = false
		default:
			if !inBracket {
				result.WriteRune(r)
			}
		}
	}
	return result.String()
}

// Transform "a" to "an" when appropriate
func TransformAToAn(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" && shouldBeAn(words[i+1]) {
			words[i] = "an"
		}
	}
	return strings.Join(words, " ")
}

func shouldBeAn(word string) bool {
	if len(word) == 0 {
		return false
	}
	first := rune(word[0])
	return strings.ContainsRune("aeiouAEIOU", first) || first == 'h' || first == 'H'
}

// Fix spacing around punctuation
func FixPunctuation(text string) string {
	output := []rune{}
	prev := rune(0)
	for i, r := range text {
		if strings.ContainsRune(".,!?:;", r) {
			if prev == ' ' {
				output = output[:len(output)-1]
			}
			output = append(output, r)
			if i < len(text)-1 && !unicode.IsPunct(rune(text[i+1])) && text[i+1] != ' ' {
				output = append(output, ' ')
			}
		} else {
			output = append(output, r)
		}
		prev = r
	}
	return string(output)
}

// Fix quotes and extra spaces
func FixQuotesAndSpaces(text string) string {
	output := []rune{}
	prev := rune(0)
	for _, r := range text {
		if r == '\'' || r == ' ' {
			if (prev == ' ' && r == '\'') || (prev == '\'' && r == ' ') {
				if prev == ' ' && r == '\'' {
					output = output[:len(output)-1]
					output = append(output, r)
				}
				continue
			}
		}
		output = append(output, r)
		prev = r
	}
	return string(output)
}
