package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Too many arguments.")
		return
	}
	// data, err := os.ReadFile(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Error reading file.")
	// 	return
	// }
	// //fmt.Println(string(data))

	// slice := strings.Split(string(data), ",")
	// for _, ch := range slice {
	// 	fmt.Println(ch)
	// }
	// fmt.Printf("Total parts: %d\n", len(slice))

	str := os.Args[1]
	// for _, ch := range str {
	// 	fmt.Printf("%c =  %d\n", ch, ch)
	// }

	for _, ch := range str {
		index := ch - ' '
		fmt.Printf("%c ascii =  %d fontIndex = %d\n", ch, ch, index)
	}

}
