package main

import "fmt"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	result := ""
	first := true

	var backtrack func(start int)

	backtrack = func(start int) {
		if len(result) == n {
			if !first {
				fmt.Print(", ")
			}
			fmt.Print(result)
			first = false
			return
		}

		for i := start; i <= 9; i++ {
			result += string(rune(i + '0'))
			backtrack(i + 1)
			result = result[:len(result)-1]
		}
	}

	backtrack(0)
}

// func main() {
// 	PrintCombN(2)
// }
