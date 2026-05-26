// Using two-pointers

package main

func TwoSum(num []int, target int) ([]int, bool) {
	newResult := []int{}
	i, j := 0, len(num)-1
	for i < j {
		if num[i]+num[j] == target {
			newResult = append(newResult, i, j)
		} else if num[i]+num[j] != target && j > i {
			j--
			continue
		}
		i++
		j--
	}
	return newResult, len(newResult) != 0
}

// func main() {
// 	fmt.Println(TwoSum([]int{-3, 5, 7, 13, 8, 9, 3, 11}, 10))
// }
