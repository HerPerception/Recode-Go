package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("empty file")
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	lines = lines[1:]
	var banner = map[rune][]string{}
	for ch := ' '; ch <= '~'; ch++ {
		start := int(ch-32) * 9
		end := start + 8
		banner[ch] = lines[start:end]
	}
	return banner, nil
}
