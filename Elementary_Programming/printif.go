package main

func printIf(s string) string {

	if len(s) == 0 {
		return "G\n"
	}

	if len(s) <= 3 {
		return "Invalid Input\n"
	}

	return "G\n"
}

// func main() {
// 	fmt.Print(printIf("abcdefz"))
// 	fmt.Print(printIf("abc"))
// 	fmt.Print(printIf(""))
// 	fmt.Print(printIf("14"))
// }
