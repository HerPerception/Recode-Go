package main

import (
	"os"
	"strings"
)

func GenerateArt(text string, filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	lines = lines[1:]
	inputSlice := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	result := []string{}
	for _, each := range inputSlice {
		for row := 0; row < 8; row++ {
			var build strings.Builder
			for _, ch := range each {
				start := int(ch-32)*9 + row
				build.WriteString(lines[start])
			}
			result = append(result, build.String())
		}
	}
	return strings.Join(result, "\n"), nil
}


