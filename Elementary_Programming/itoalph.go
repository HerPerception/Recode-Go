package main

func Itoa(n int) string {
	str := ""
	s := ""
	if n == 0 {
		return string(rune(n) + '0')
	}
	if n < 0 {
		n = n * -1
		for n > 0 {
			digit := n % 10
			s = string(rune(digit)+'0') + s
			n = n / 10
		}
		str = "-" + s
	}
	if n > 0 {
		for n > 0 {
			digit := n % 10
			str = string(rune(digit)+'0') + str
			n = n / 10
		}
	}
	return str
}

// func main() {
// 	fmt.Println(Itoa(12345))
// 	fmt.Println(Itoa(0))
// 	fmt.Println(Itoa(-1234))
// 	fmt.Println(Itoa(987654321))
// }
