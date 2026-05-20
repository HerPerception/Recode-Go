package main

import "fmt"

func main() {

	for i := 9; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			for k := j - 1; k >= 0; k-- {
				if i == 2 && j == 1 && k == 0 {
					fmt.Printf("%d%d%d", i, j, k)
				} else {
					fmt.Printf("%d%d%d, ", i, j, k)
				}
			}
		}
	}
}
