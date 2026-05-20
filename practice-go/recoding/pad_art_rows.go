package main

import "strings"

func PadArtRows(rows []string, width int) []string {
	result := make([]string, len(rows))
	for i, row := range rows {
		padding := width - len(row)
		if padding > 0 {
			result[i] = row + strings.Repeat(" ", padding)
		}
		if padding <= 0 {
			result[i] = row
		}
	}
	return result
}
