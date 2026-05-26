package main

func rectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}
	return 2 * (w + h)
}

// func main() {
// 	fmt.Println(rectPerimeter(10, 2))
// 	fmt.Println(rectPerimeter(434343, 898989))
// 	fmt.Println(rectPerimeter(10, -2))
// }
