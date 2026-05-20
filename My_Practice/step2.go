package main

import "fmt"

func main() {
	sInt := []int{2, 3, 4, 5, 9, 12}
	sum, min, max, even, odd, average, above, below, equalTo := sumSlice(sInt)
	fmt.Printf("The sum of the slice is %d, and the minimum value is %d, while the maximum value is %d, even numbers count is %d and odd numbers count is %d. The average is %.2f and the values above average are %v and the values below average are %v, the value equal to average is %v.\n", sum, min, max, even, odd, average, above, below, equalTo)
}

func sumSlice(sInt []int) (sum int, min int, max int, even int, odd int, average float64, above []int, below []int, equalTo []int) {
	if len(sInt) == 0 {
		return
	}
	sum = 0
	min = sInt[0]
	max = sInt[0]
	above = []int{}
	even = 0
	odd = 0

	for _, i := range sInt {
		sum += i
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
		if i%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	average = float64(sum) / float64(len(sInt))
	for _, i := range sInt {
		if float64(i) > average {
			above = append(above, i)
		} else if float64(i) < average {
			below = append(below, i)
		} else {
			equalTo = append(equalTo, i)
		}
	}
	return sum, min, max, even, odd, average, above, below, equalTo
}
