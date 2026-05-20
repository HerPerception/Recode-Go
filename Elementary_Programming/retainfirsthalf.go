package main

import "fmt"

func retainFirstHalf(s string)string {
	if len(s) == 0{
		return ""
	}
	if len(s) == 1 {
		return s
	}
	return s[:(len(s)/2)]
}

func main() {
	fmt.Println(retainFirstHalf("This is the 1st HalfThis is the 2nd half"))
	fmt.Println(retainFirstHalf("A"))
	fmt.Println(retainFirstHalf(""))
	fmt.Println(retainFirstHalf("Hello World"))
}