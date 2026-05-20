package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(i int) {
	str := ""
	neg := false
	if i < 0 {
		neg = true
		i *= -1
	} else if i == 0 {
		z01.PrintRune('0')
	}
	for i > 0 {
		digit := i % 10
		str = string(rune(digit)+'0') + str
		i = i / 10
	}
	if neg {
		z01.PrintRune('-')
	}
	for _, char := range str {
		z01.PrintRune(char)
	}
}
func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
