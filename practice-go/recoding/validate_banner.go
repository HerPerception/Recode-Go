package main

import (
	"errors"
	"fmt"
)

func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return errors.New("empty banner map")
	}
	if len(banner) != 95 {
		return errors.New("invalid banner map")
	}

	for r := rune(32); r <= rune(126); r++ {
		if len(banner[r]) != 8 {
			return fmt.Errorf("character %c has %d expected 8", r, len(banner[r]))
		}
	}
	return nil
}
