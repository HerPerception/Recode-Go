package main

func StackTwo(top []string, bottom []string) []string {
	if len(top) == 0 && len(bottom) == 0 {
		return []string{}
	}
	return append(top, bottom...)
}

func StackAll(blocks [][]string) []string {
	result := []string{}
	for _, block := range blocks {
		result = StackTwo(result, block)
	}
	return result
}
