package main

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	words := []string{}
	word := ""
	for _, i := range s {

		if i != ' ' {
			word += string(i)
		}
		if i == ' ' && word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}

	for i := range words {
		str := words[i]
		if str[0] >= 'a' && str[0] <= 'z' {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsCapitalized("Hello! How are you?"))
// 	fmt.Println(IsCapitalized("Hello How Are You"))
// 	fmt.Println(IsCapitalized("Whats 4this 100K?"))
// 	fmt.Println(IsCapitalized("Whatsthis4"))
// 	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
// 	fmt.Println(IsCapitalized(""))
// }
