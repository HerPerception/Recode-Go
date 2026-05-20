package main

import (
	"strconv"
	"strings"
)

func Case(s string) string {
	str := strings.Fields(s)
	num := 1
	for i := 0; i < len(str); i++ {
		if strings.HasPrefix(str[i], "(") {
			if strings.HasSuffix(str[i+1], ")") {
				temp, err := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				if err != nil {
					return s
				} else {
					num = temp
				}
			} else {
				num = 1
			}
		}
		if num > i {
			num = i
		}
		for j := i - num; j < i; j++ {
			switch str[i] {
			case "(up,", "(up)":
				str[j] = strings.ToUpper(str[j])

			case "(low,", "(low)":
				str[j] = strings.ToLower(str[j])

			case "(cap,", "(cap)":
				str[j] = strings.ToUpper(string(str[j][0])) + strings.ToLower(str[j][1:])
			}
		}
		if strings.HasPrefix(str[i], "(") {
			str = append(str[:i], str[i+1:]...)
			i--
		} else if strings.HasSuffix(str[i], ")") {
			str = append(str[:i], str[i+1:]...)
			i--
		}
	}
	return strings.Join(str, " ")
}
