package main

import "fmt"

type StatsResult struct {
	Sum     int
	Min     int
	Max     int
	Average float64
}

func Stats(slice []int) (StatsResult, error) {

	if len(slice) == 0 {
		return StatsResult{}, fmt.Errorf("slice is empty")
	}
	max := slice[0]
	min := slice[0]
	sum := 0
	for _, i := range slice {
		sum += i
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	average := float64(sum) / float64(len(slice))
	result := StatsResult{
		Sum:     sum,
		Min:     min,
		Max:     max,
		Average: average,
	}
	return result, nil
}

func main() {

	fmt.Println(Stats([]int{6, 7, 8, 9}))
}
