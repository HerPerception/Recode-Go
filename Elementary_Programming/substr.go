package main

func Substr(str string, substr string) int {
	for i := range str {
		if len(str) == len(substr) && str == substr {
			return 0
		}
		if len(substr) == 1 && string(str[i]) == substr {
			return i
		}
		if i+len(substr) <= len(str) && str[i] == substr[0] && str[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
func strStr(str string, substr string) int {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:len(substr)+i] == substr {
			return i
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(strStr("sYdbutsad", "sad"))
// }
