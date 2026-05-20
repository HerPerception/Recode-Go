package main

import "fmt"

type StatsResult struct {
	Sum int
	Min int
	Max int
    Average float64
}

func Stats(slice []int) (map[string]float64, error) {
	s := StatsResult{
		Sum: 0
		Min: slice[0]
		Max: slice[0]
        Average: 
	}
	if len(slice) == 0 {
		return nil, fmt.Errorf("slice is empty")
	}
	result := map[string]float64{}
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
	result["sum"] = float64(sum)
	result["max"] = float64(max)
	result["min"] = float64(min)
	result["average"] = float64(sum) / float64(len(slice))
	return result, nil
}

func main() {

	fmt.Println(Stats([]int{6, 7, 8, 9}))
}
