package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CmdN(s []string) []string {
	for i := 0; i < len(s); i++ {
		num := 0
		cmd := ""
		if strings.HasPrefix(s[i], "(") && strings.HasSuffix(s[i+1], ")") {
			cmd = s[i][1:]
			temp, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])
			if err == nil {
				num = temp
			}
		}
		if num > i {
			num = i
		}
		for j := i - num; j < i; j++ {
			if cmd == "up" {
				s[j] = strings.ToUpper(s[j])
			} else if cmd == "low" {
				s[j] = strings.ToLower(s[j])
			} else if cmd == "cap" {
				s[j] = strings.ToUpper(string(s[j][0])) + strings.ToLower(s[j][1:])
			}
		}
		if strings.HasPrefix(s[i], "(") {
			s = append(s[:i], s[i+1:]...)
			i--
		} else if strings.HasSuffix(s[i], ")") {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func main() {
	fmt.Println(CmdN([]string{"my", "mum", "(up", "6)", "iS", "VERY", "NICE", "(low", "3)", "and", "my", "dad", "too", "(cap", "4)"}))
}
