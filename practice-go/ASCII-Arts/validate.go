package main

import "errors"

func ValidateInput(input string) (rune, error) {
	for _, ch := range input {
		if ch < 32 || ch > 126 {
			return ch, errors.New("Unsupported character")
		}
	}
	return 0, nil
}
