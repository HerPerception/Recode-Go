package main

import "fmt"

func itoa(i int)string {
	str := ""
	for i > 0 {
		digit := i%10
		str = string(rune(digit) + '0') + str
		i = i/10
	}
	return str
}

func atoi(i string) int {
	num := 0
	for _, idx := range i {
		digit := idx -'0'
		num = num*10 + int(digit)
	}
	return num
}
func main(){
	fmt.Printf("%q\n", itoa(12345))
	fmt.Printf("%d\n", atoi("12345"))
}