package main

import "fmt"

// func FirstDuplicate(slice []any) (any, bool) {
// 	myMap := make(map[any]string)
// 	var val any
// 	for _, r := range slice {
// 		_, key := myMap[r]
// 		if key {
// 			val = r
// 			return val, true
// 		} else {
// 			myMap[r] = "there"
// 		}
// 	}
// 	return val, false
// }

func FirstDuplicate[T comparable](items []T) (T, bool) {
	myMap := make(map[T]string)
	var val T
	for _, r := range items {
		_, exists := myMap[r]
		if exists {
			val = r
			return val, true
		} else {
			myMap[r] = "there"
		}
	}

	return val, false
}
func main() {
	fmt.Println(FirstDuplicate([]int{1, 2, 3, 2, 1}))
	fmt.Println(FirstDuplicate([]string{"go", "js", "go", "js"}))
	fmt.Println(FirstDuplicate([]int{}))
	fmt.Println(FirstDuplicate([]int{1, 2, 3}))
	fmt.Println(FirstDuplicate([]string{"a", "b", "c"}))
	fmt.Println(FirstDuplicate([]bool{false, true, false}))
}
