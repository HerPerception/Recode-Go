package main

//	func moveZeros(inputSlice []int) []int {
//		num := []int{}
//		count := 0
//		for _, ch := range inputSlice {
//			if ch == 0 {
//				count++
//			} else {
//				num = append(num, ch)
//			}
//		}
//		for range count {
//			num = append(num, 0)
//		}
//		return num
//	}

func moveZeros(inputSlice []int) []int {
	left, right := 0, 0
	for right < len(inputSlice) {
		if inputSlice[right] == 0 {
			right++
			continue
		}
		if inputSlice[right] != 0 {
			temp := inputSlice[left]
			inputSlice[left] = inputSlice[right]
			inputSlice[right] = temp
			left++
		}
		right++

	}
	return inputSlice
}

// func main() {
// 	fmt.Println(moveZeros([]int{0, 2, 2, 0, 3, 4, 0, 3, 0, 2, 0}))
// 	fmt.Println(moveZeros([]int{1, -2, 12, 0, 4, 5, 0, 5, 4, 6}))
// }
