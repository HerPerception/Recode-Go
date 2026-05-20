package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	first, second := os.Args[1], os.Args[2]

	a := 0
	for _, i := range first {
		digit := i -'0'
		a = a * 10 + int(digit)
	}
	b := 0
	for _, i := range second {
		digit := i -'0'
		b = b * 10 + int(digit)
	}
	if b > a {
		temp := a 
		a = b
		b = temp
	}
	c := a%b
	for c > 0 {
		a = b
		b = c
		c = a%b
	}
	fmt.Println(b)
}
