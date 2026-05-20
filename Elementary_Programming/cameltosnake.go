package main

import "fmt"

func camelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	isLetter := true

	for index, ch := range s {
		if (ch < 'A' && ch > 'Z') && (ch < 'A' && ch > 'z') {
			isLetter = false
		}
		if index == len(s)-1 && (ch >= 'A' && ch <= 'Z') {
			isLetter = false
		}
		if index < len(s)-1 && (s[index] >= 'A' && s[index] <= 'Z') && (s[index+1] >= 'A' && s[index+1] <= 'Z') {
			isLetter = false
		}
	}
	word := ""
	if isLetter == false {
		return s
	}
	for index, ch := range s {
		if index != 0 && (ch >= 'A' && ch <= 'Z') {
			word += "_" + string(ch)
		} else {
			word += string(ch)
		}
	}
	return word
}

func main() {
	fmt.Println(camelToSnakeCase("HellWorld"))
	fmt.Println(camelToSnakeCase("helLOWORld"))
	fmt.Println(camelToSnakeCase("camelCase"))
	fmt.Println(camelToSnakeCase("CamelCase"))
	fmt.Println(camelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(camelToSnakeCase("CamelToSnakeCase"))
	fmt.Println(camelToSnakeCase("hey2"))

}
