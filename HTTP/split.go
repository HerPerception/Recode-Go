package main

import (
	"strings"
)

func SplitInput(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
