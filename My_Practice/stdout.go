package main

import (
	"fmt"
)

func main() {
	/*if len(os.Args) < 3 {
		fmt.Println("Please provide complete  files.")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file.")
	}
	err = os.WriteFile(outputFile, file, 0644)
	if err != nil {
		fmt.Println("Error writing file.")
	}
	fmt.Println("Process completed.")

	for i := 1; i <= 20; i = i + 2 {
		fmt.Println(i)
	} */
	fmt.Println(indexOfPunc("Ebele."))
	fmt.Println(intToString(15))
	fmt.Printf("%q\n", convertToString(20))
	fmt.Printf("%d\n", convertToInt("15"))
}
func countLetters(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {
			count++
		}
	}
	return count
}

func isLetters(s string) bool {
	for _, val := range s {
		if val < 'A' || val > 'z' {
			return false
		}
	}
	return true
}

func indexOfPunc(s string) (int, string) {
	for index, val := range s {
		if val < 'A' || val > 'z' {
			return index, string(val)
		}
	}
	return 0, s
}

func intToString(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	for n > 0 { // <-- loop starts
		digit := n % 10                           // get last digit (remainder)
		result = string(rune(digit)+'0') + result // prepend digit to result
		n /= 10                                   // remove last digit
	} // <-- loop ends

	return result
}

func convertToString(i int) string {
	str := ""
	for i > 0 {
		digit := i % 10
		str = string(rune(digit)+'0') + str
		i = i / 10
	}
	return str
}

func convertToInt(s string) int {
	num := 0
	for _, i := range s {
		digit := i - '0'
		num = num*10 + int(digit)
	}
	return num
}
