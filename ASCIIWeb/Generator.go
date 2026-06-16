package main

import (
	"fmt"
	"os"
	"strings"
)

func Generator(input, filename string) (string, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	banner := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	banner = banner[1:]
	inputSlice := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	var result []string
	for _, line := range inputSlice {
		for row := 0; row < 8; row++ {
			var build strings.Builder
			for _, each := range line {
				if each < 32 || each > 126 {
					return "", fmt.Errorf("%c is an unsupported character.", each)
				}
				start := int(each-32)*9 + row
				build.WriteString(banner[start])
			}
			result = append(result, build.String())
		}

	}
	return strings.Join(result, "\n"), nil
}
