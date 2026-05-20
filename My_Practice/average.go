package main

import "fmt"

func main() {
	fmt.Println(Average)
}

func Average(numbers []float64) float64 {
	sum := 0.0
	count := 0.0
	fmt.Println("Input the numbers here:")
	fmt.Scanf("%v", &numbers)
	for _, i := range numbers {
		sum += i
		count++
	}
	average := sum / count
	return average
}
