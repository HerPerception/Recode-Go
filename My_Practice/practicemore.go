package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWORld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	isCamel := true
	for i := range s {
		if s[i] < 'a' && s[i] > 'z' || s[i] < 'A' && s[i] > 'Z' {
			isCamel = false
		}
		if i < len(s)-1 && (s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
			isCamel = false
		}
		if i == len(s)-1 && (s[i] >= 'A' && s[i] <= 'Z') {
			isCamel = false
		}
	}
	word := ""
	if isCamel == true {
		for i := range s {

			if (i > 0 && i < len(s)-1) && (s[i] >= 'A' && s[i] <= 'Z') {
				word += "_" + string(s[i])
			} else {
				word += string(s[i])
			}
		}
		return word
	}
	return s

}
