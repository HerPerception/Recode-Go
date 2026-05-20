package main

import (
	"fmt"
	"os"
)

func main() {
	myFile, _ := os.Create("output.txt")
	defer myFile.Close()
	fmt.Fprint(myFile, "This is a new way to write to a file. I learned it today.")
}
