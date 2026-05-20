package main

import "fmt"

func main() {
	num := 0
	number := 0
	operator := ""
	fmt.Print("Enter number: ")
	fmt.Scanln(&num)
	fmt.Print("Enter operator: ")
	fmt.Scanln(&operator)
	if operator != "+" {
		fmt.Println("invalid operator")
	} else {
		fmt.Println("Enter second number: ")
	}
	fmt.Scanln(&number)
	fmt.Println(num + number)
}
