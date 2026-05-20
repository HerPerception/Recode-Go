package main

import (
	//"strings"
	"errors"
	//"errors"
	"fmt"
)

func ValidateBanner(banner map[rune][]string) error {
	if len(banner) != 95 {
		return errors.New("empty banner file")
	}
	for key, value := range banner {
		if key < 32 || key > 126 {
			return fmt.Errorf("unsupported characters")
		}
		if len(value) != 8 {
			return fmt.Errorf("invalid lines")
		}
	}
	return nil
}
