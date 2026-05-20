package main

import "fmt"

func StatsResults(num []int) (sum int, average float64, min int, max int, count int, counter int) {
	if len(num) == 0 {
		fmt.Println("Slice is empty, sum is 0")
		return 0, 0, 0, 0, 0, 0
	}
	sum = 0
	min = num[0]
	max = num[0]
	count = 0
	counter = 0
	for _, i := range num {
		fmt.Println("Adding", i)
		sum += i
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
		if i%2 == 0 {
			count++
		}
		if i%2 != 0 {
			counter++
		}
	}
	average = float64(sum) / float64(len(num))
	return sum, average, min, max, count, counter
}

func main() {
	num := []int{8, 2, 5, 6}
	sum, average, min, max, count, counter := Average(num)
	fmt.Printf("Sum: %d, Average: %.1f, min: %d, max: %d, Even numbers count: %d, Odd numbers count: %d\n", sum, average, min, max, count, counter)
}

/* Adding 8
Adding 2
Adding 5 */
