package main

import "fmt"

func main() {
	fmt.Println(GCD("3", "17"))
}

func GCD(s, str string) int {
	var digit int
	num := 0
	for i := range s {
		num *= 10
		digit = int(s[i] - '0')
		num += digit
	}
	number := 0
	for i := range str {
		number *= 10
		digit = int(str[i] - '0')
		number += digit

	}
	if num < number {
		temp := num
		num = number
		number = temp
	}
	c := num % number
	for c > 0 {
		num = number
		number = c
		c = num % number
	}
	return number
}
