package main

import "fmt"

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k <= 9; k++ {
				for l := j + 1; l <= 9; l++ {
					num := i*10 + j
					num2 := k*10 + l
					if num > num2 {
						continue
					} else if i == 9 && j == i-1 && k == 9 {
						fmt.Printf("%d%d %d%d\n", i, j, k, l)
					} else {
						fmt.Printf("%d%d %d%d, ", i, j, k, l)
					}
				}
			}
		}
	}
}

func main() {
	PrintComb2()
}
