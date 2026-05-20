package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	result := make(map[rune][]string)
	for key, val := range base {
		result[key] = val
	}
	for key, val := range priority {
		result[key] = val
	}
	return result
}
