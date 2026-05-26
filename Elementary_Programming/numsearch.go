package main

func search(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(search([]int{4, 5, 6, 7, 3, 2, 1, 0}, 7))
// }
