package main

import "fmt"

func fishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	} else if n%6 == 0 {
		return "fish and chips"
	} else if n%2 == 0 {
		return "fish"
	} else if n%3 == 0 {
		return "chips"
	}
		return "error: non-divisible"
}
func main(){
	fmt.Println(fishAndChips(4))
	fmt.Println(fishAndChips(9))
	fmt.Println(fishAndChips(6))
	fmt.Println(fishAndChips(-4))
	fmt.Println(fishAndChips(7))
}