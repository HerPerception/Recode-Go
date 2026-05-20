package main

import "fmt"

func main() {
	j := 'z'
	i := 'a'
	for i <= j {
		for j >= i {
			fmt.Print(string(i), (string(j)))

			i++
			j--
		}
	}

}
