package main

import "fmt"

func main() {
	M := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
	}
	for _, value := range M {
		fmt.Println(value)
	}
}
