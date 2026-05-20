package main

import "fmt"

func main() {
	//var name string = "Osinaka"
	//age := 22
	//fmt.Printf("Hello, I'm %v and I'm %v years old.\n", name, age)
	//fmt.Println("Hello Go!")

	/*nums := []int{2, 4, 7, 9, 5}
	for _, i := range nums {
		fmt.Println(i)
	}*/
	fmt.Println(Sum(4, 9))
}

func Sum(i, n int) int {
	return i + n
}
