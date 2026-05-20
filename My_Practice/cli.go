package main

import "fmt"

func main() {
	for {
		fmt.Print("Enter first number: ")
		var first int
		fmt.Scanln(&first)
		fmt.Print("Enter operator(+, -, *, /, %, percent): ")
		var operator string = ""
		var second int
		fmt.Scanln(&operator)
		if operator != "+" && operator != "-" && operator != "*" && operator != "/" && operator != "%" && operator != "percent" {
			fmt.Println("Invalid operator")
			return
		} else {
			fmt.Print("Enter second number: ")
		}
		fmt.Scanln(&second)

		switch operator {
		case "+":
			fmt.Printf("%d + %d = %d", first, second, first+second)
			fmt.Println()

		case "-":
			fmt.Printf("%d - %d = %d", first, second, first-second)
			fmt.Println()

		case "*":
			fmt.Printf("%d * %d n = %d", first, second, first*second)
			fmt.Println()

		case "/":
			if second == 0 {
				fmt.Println("No division by zero")
			}
			if first%second != 0 {
				fmt.Printf("%d / %d = %.2f", first, second, first/second)
				fmt.Println()
			} else {
				fmt.Printf("%d / %d = %d", first, second, first/second)
				fmt.Println()
			}
		case "%":
			fmt.Printf("%d divided by %d would leave %d as remainder", first, second, first%second)
		case "percent":
			fmt.Printf("%d%% of %d = %d", first, second, (first * second / 100))
			fmt.Println()
		}
		var options string
		fmt.Print("Choose options(continue, exit): ")
		fmt.Scanln(&options)
		if options == "exit" {
			break
		}
	}
}
