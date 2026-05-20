package mypractice
package main

import (
	"fmt"
	"strings"
)

func UpperLastWords(s []string) []string {
	for i := range s {
		s[i] = strings.ToLower(s[i])
	}
	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s = append(s[:i], s[i+1:]...)
		} else if s[i] == "(low)" {
			s[i-1] = strings.ToLower(s[i-1])
			s = append(s[:i], s[i+1:]...)
		} else if s[i] == "(cap)" {
			s[i-1] = strings.ToUpper(string(s[i-1][0])) + strings.ToLower(s[i-1][1:])
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
func main() {
	s := []string{"yes", "recoding", "has", "begun", "go", "(up)", "take", "(cap)", "things", "(cap)", "very", "(cap)", "cAREfully", "(cap)", "EASY", "(low)"}
	fmt.Println(UpperLastWords(s))
}
