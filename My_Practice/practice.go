package main

import "github.com/01-edu/z01"

func main() {
	j := 'z'
	i := 'a'
	for i <= j {
		for j >= i {
			z01.PrintRune(i)
			z01.PrintRune(j)
			i++
			j--
		}
	}
	
	
}
