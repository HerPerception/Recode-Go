package main

import (
	"fmt"
)

//	func miniMaxSum(arr []int32){
//		res := []int64{}
//		sum := int64(0)
//		for _, num := range arr{
//			sum += int64(num)
//		}
//		for _, num := range arr {
//			x := sum - int64(num)
//			res = append(res, int64(x))
//		}
//		min := res[0]
//		max := res[0]
//		for _, k := range res {
//			if k > max {
//				max = k
//			}
//			if k < min {
//				min = k
//			}
//		}
//		fmt.Println(min, max)
//	}

// func main() {
// 	miniMaxSum([]int32{1, 2, 3, 4, 5})
// }

func miniMaxSum(arr []int32) {
	sum := int64(0)
	min := int64(0)
	max := int64(0)
	for _, num := range arr {
		sum += int64(num)
	}
	for _, num := range arr {
		x := int64(0)
		x = sum - int64(num)
		min = x
		if max < x {
			max = x
		}
		if min > x {
			min = int64(x)
		}
	}
	fmt.Println(min, max)
}
