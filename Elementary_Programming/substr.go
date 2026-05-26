package main

func Substr(str string, substr string)int {
	for i := range str {
		if i+len(substr) < len(str) && str[0] == substr[0] && str[i:i+len(substr)] == substr{
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(Substr("a", "a"))
}