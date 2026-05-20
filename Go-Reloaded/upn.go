package main

import (
	"fmt"
	"strings"
)

func UpN(s []string) []string {
	cmd := ""
	num := 1
	//temp := 0
	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i], "(") && strings.HasSuffix(s[i+1], ")") {
			cmd = s[i][1:]
			num = int(s[i+1][0] - '0')
			//temp = i - num
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
	fmt.Println(UpN([]string{"my", "mum", "(up", "6)", "iS", "VERY", "NICE", "(low", "4)", "and", "my", "dad", "too", "(cap", "4)"}))
}
