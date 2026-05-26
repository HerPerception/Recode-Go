package main

import "fmt"

func Substr(str string, substr string) int {
	for i := range str {
		if len(str) == len(substr) && str == substr {
			return 0
		}
		if len(substr) == 1 && string(str[i]) == substr {
			return i
		}
		if i+len(substr) < len(str) && str[0] == substr[0] && str[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(Substr("ab", "ab"))
}
