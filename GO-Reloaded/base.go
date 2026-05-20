package main

import (
	"strconv"
	"strings"
)

func Base(s string) string {
	str := strings.Fields(s)
	for i := 1; i < len(str); i++ {
		switch str[i] {
		case "(hex)":
			temp, err := strconv.ParseInt(str[i-1], 16, 64)
			if err != nil {
				return s
			}
			str[i-1] = strconv.FormatInt(temp, 10)
			str = append(str[:i], str[i+1:]...)
			i--

		case "(bin)":
			temp, err := strconv.ParseInt(str[i-1], 2, 64)
			if err != nil {
				return s
			}
			str[i-1] = strconv.FormatInt(temp, 10)
			str = append(str[:i], str[i+1:]...)
			i--
		}
	}
	return strings.Join(str, " ")
}
