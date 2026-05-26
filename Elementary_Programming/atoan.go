package main

func aToAn(s string) string {
	if s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u' || s[0] == 'h' {
		return "an " + s
	}
	return "a " + s
}

func Vowel(slice []string) string {
	// slice := []string{}
	// for _, ch := range s {
	// 	slice = append(slice, string(ch))
	//}
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == "a" && (slice[i+1][0] == 'a' || slice[i][0] == 'e' || slice[i][0] == 'i' || slice[i][0] == 'o' || slice[i][0] == 'u' || slice[i][0] == 'h') {
			slice[i] = "an"
		} else if slice[i] == "A" && (slice[i+1][0] == 'a' || slice[i][0] == 'e' || slice[i][0] == 'i' || slice[i][0] == 'o' || slice[i][0] == 'u' || slice[i][0] == 'h') {
			slice[i] = "An"
		}
	}
	word := ""
	for _, ch := range slice {
		word += string(ch)
	}
	return word
}

// func main() {
// 	fmt.Println(aToAn("book"))
// 	fmt.Println(aToAn("apple"))
// 	fmt.Println(Vowel([]string{"There", "it", "was", ".", "A", "amazing", "rock", "under", "a", "orange", "tree", ".", "A", "book"}))
// }
