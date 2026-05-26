package main

func printIfNot(s string) string {
	if len(s) == 0 {
		return "G\n"
	}
	if len(s) < 3 {
		return "G\n"
	}
	return "Invalid input\n"
}

// func main() {
// 	fmt.Print(printIfNot("abcdefz"))
// 	fmt.Print(printIfNot("abc"))
// 	fmt.Print(printIfNot(""))
// 	fmt.Print(printIfNot("14"))
// }
