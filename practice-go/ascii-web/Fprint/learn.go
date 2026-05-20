package main

import (
	"fmt"
	"os"
)

func main() {
	filename, _ := os.Create("output.txt")
	defer filename.Close()
	fmt.Fprint(filename, "I am learning Go for fun and I love it!")
	fmt.Fprint(filename, "I am learning Go for fun and I love it!")
}
