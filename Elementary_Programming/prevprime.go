package main

import "fmt"

func FindPrevPrime(nb int) int {
	n := 0
	prime := true
	for i := nb; i >= 2; i-- {
		prime = true
		for n := 2; n*n <= i; n++ {
			if i%n == 0 {
				prime = false
				break
			}

		}
		//prime = true
		if prime != false && i <= nb {
			n = i
			break
		}
	}
	return n
}

func main() {
	fmt.Println(FindPrevPrime(34))
}
