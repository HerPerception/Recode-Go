package main

import (
	"fmt"
)

func main() {
	arr1 := []int{6, 4, 2, 1, 90, -1}
	sort.(arr1)
	last_num := arr1[len(arr1)-1]
	fmt.Println(last_num)
}
