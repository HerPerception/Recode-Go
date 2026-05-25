//Solving this question using map

package main

import "fmt"

func TwoSum(num []int, target int) ([]int, bool) {
	myMap := make(map[int]int)
	newResult := []int{}
	for i, each := range num {
		number := target - each
		val, ok := myMap[number]
		if ok {
			newResult = append(newResult, val, i)
		} else {
			myMap[each] = i
		}
	}

	return newResult, len(newResult) != 0
}

func main() {
	fmt.Println(TwoSum([]int{-3, 5, 7, 8, 9, 3, 11, 13}, 1))
}
