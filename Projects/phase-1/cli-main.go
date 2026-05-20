package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	fmt.Println("Welcome To CLI Calculator")
CalculatorLoop:
	for {
		fmt.Println("Choose operation: (Calculator, Help)")
		scanner.Scan()
		input = scanner.Text()
		if input == "Help" {
			fmt.Println("Addition: num1 + num2")
			fmt.Println("Subtraction: num1 - num2")
			fmt.Println("Multiplication: num1*num3")
			fmt.Println("Division: num1/num2")
		}

		if input == "Calculator" {
			fmt.Println("Enter values to perform operation:")
			scanner.Scan()
			newInput := scanner.Text()
			var newStr string
			//var flag int
			for _, ch := range newInput {
				if ch == ' ' {
					continue
				} else {
					newStr += string(ch)
				}
			}
			split := 0
			operator := ""
			for i := range newStr {
				if i > 0 && (newStr[i] == '+' || newStr[i] == '-' || newStr[i] == '*' || newStr[i] == '/') {
					operator = string(newStr[i])
					split = i
					break
				}

			}
			if len(newStr[:split]) == 0 || len(newStr[split+1:]) == 0 {
				fmt.Println("Wrong number of arguments. You must provide two numbers. Type Help to see how.")
				continue CalculatorLoop
			}
			firstNum := 0
			secondNum := 0
			first, err := strconv.Atoi(newStr[:split])
			if err != nil {
				fmt.Println("Error. Check input and try again. Type (Help) to see how the operation works.")
				continue CalculatorLoop
			} else {
				firstNum = first
			}

			second, err := strconv.Atoi(newStr[split+1:])
			if err != nil {
				fmt.Println("Error. Check input and try again. Type (Help) to see how the operation works.")
				continue CalculatorLoop
			} else {
				secondNum = second
			}

			switch operator {
			case "+":

				fmt.Printf("%d + %d = %d\n", firstNum, secondNum, firstNum+secondNum)

			case "-":
				fmt.Printf("%d - %d = %d\n", firstNum, secondNum, firstNum-secondNum)

			case "*":
				fmt.Printf("%d * %d = %d\n", firstNum, secondNum, firstNum*secondNum)
			case "/":

				if secondNum == 0 {
					fmt.Println("No division by 0")
				} else {
					if firstNum%secondNum != 0 {
						fmt.Printf("%d / %d = %.3f\n", firstNum, secondNum, float64(firstNum)/float64(secondNum))
					} else {
						fmt.Printf("%d / %d = %d\n", firstNum, secondNum, firstNum/secondNum)
					}
				}
			}
		}

		fmt.Println("Choose option: (continue, exit)")
		scanner.Scan()
		option := scanner.Text()
		if option == "exit" {
			fmt.Println("Goodbye from CLI Calculator.")
			break
		} else {
			continue
		}

	}
}
