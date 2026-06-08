package main

import (
	"errors"
	"os"
)

func MakeBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("empty file")
	}
	line := SplitInput(string(data))
	line = line[1:]
	if len(line) != 855 {
		return nil, errors.New("invalid file content")
	}
	banner := map[rune][]string{}
	for i := 32; i <= 126; i++ {
		start := int(i-32) * 9
		end := start + 8
		banner[rune(i)] = line[start:end]
	}
	return banner, nil
}
