package main

func digitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n = n * (-1)
	}
	if n == 0 {
		return 1
	}
	count := 0

	for n > 0 {
		n = n / base
		count++
	}
	return count
}

// func main() {
// 	fmt.Println(digitLen(100, 10))
// 	fmt.Println(digitLen(100, 2))
// 	fmt.Println(digitLen(-100, 16))
// 	fmt.Println(digitLen(100, -1))
// }
