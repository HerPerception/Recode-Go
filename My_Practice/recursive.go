package main

import "fmt"

func recursiveSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + recursiveSum(numbers[1:])
}
func main() {
	nums := []int{1, 2, 3, 4, -4, -7, 10, -12}
	fmt.Println(recursiveSum(nums))
}
