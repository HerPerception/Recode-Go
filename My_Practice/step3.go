package main

import "fmt"

func main() {
	data := map[string]int{"average": 20, "sum": 200,
		"mean": 2}
	data["age"] = 30
	//delete(data, "age")
	value, ok := data["fire"]
	fmt.Println(value, ok)
}
