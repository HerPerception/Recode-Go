package main

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	var run func(start int, comb []byte)

	run = func(start int, comb []byte) {
		if len(comb) == n {
			for i := 0; i < n; i++ {
				z01.PrintRune(rune(comb[i] + '0'))
			}
			if comb[0] != byte(10-n) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}

		for i := start; i <= 9; i++ {
			run(i+1, append(comb, byte(i)))
		}
	}

	run(0, []byte{})
}

func main() {
	PrintCombN(5)
}
