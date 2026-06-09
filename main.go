package main

import (
	"fmt"
)

func Build[T comparable](anyData []T, newItem T) []T {
	return append(anyData, newItem)
}
func main() {
	word := []map[interface{}]interface{}{
		{
			"a": "apple",
		},

		{
			"number": 800,
		},
		{
			"ans": true,
		},

		{
			"height": 1.9876,
		},
		{
			89: false,
		},
	}

	fmt.Println(word[1])
	fmt.Println(word)

	result := Build([]string{}, "yesterday")
	result2 := Build([]int{1, 3, 4}, 2)
	result3 := Build([]any{2, "hello", false}, 3.45)
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	

}
