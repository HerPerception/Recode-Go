package main

import (
	"errors"
	"fmt"
)

func Validate(text string) (rune, error) {
	for _, ch := range text {
		if ch < 32 || ch > 126 {
			err := fmt.Sprintf("unsupported character: %c", ch)
			return ch, errors.New(err)
		}
	}
	return 0, nil
}
