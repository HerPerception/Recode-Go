package main

func RenderLine(input string, banner map[rune][]string) []string {
	var str []string
	for row := 0; row < 8; row++ {
		word := ""
		for _, ch := range input {
			word += banner[ch][row]
		}
		str = append(str, word)
	}
	return str
}
