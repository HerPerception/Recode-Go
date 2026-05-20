package main

import "fmt"

func hashCode(dec string)string {
	hash := ""
	for _, ch := range dec {
		if (int(ch) + len(dec)% 127) < 33 {
			hash += string(ch + 33)
		} else {
			hash += string(int(ch) + len(dec)% 127)
		}
	}
	return hash
}

func main(){
	fmt.Println(hashCode("A"))
	fmt.Println(hashCode("AB"))
	fmt.Println(hashCode("BAC"))
	fmt.Println(hashCode("Hello World"))
}