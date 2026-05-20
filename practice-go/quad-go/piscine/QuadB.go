package piscine

import "fmt"

func QuadB(x, y int) {
	for i := 0; i <= y-1; i++ {
		for j := 0; j <= x-1; j++ {
			if j != 0 && j != x-1 && (i == 0 || i == y-1) {
				fmt.Print("*")
			} else if (i > 0 && i <= y-2) && (j == 0 || j == x-1) {
				fmt.Print("*")
			} else if (i > 0 && i <= y-2) && (j != 0 || j != x-1) {
				fmt.Print(" ")
			} else if (i == 0 && j == 0) || (i == y-1 && j == x-1) {
				fmt.Print("/")
			} else if (i == 0 && j == x-1) || (i == y-1 || j == 0) {
				fmt.Print("\\")
			}
		}
		fmt.Println()
	}
}

// if j != 0 && j != x-1 && (i == 0 || i == y-1)
//  else if (j == 0 || j == x-1) && (i == 0 || i == y-1)

// func main() {
// 	QuadB(1, 5)
// }
