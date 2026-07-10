package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
guessStart:
	for {
		fmt.Println("I'm thinking of a number between 1 and 100... Can you guess what number it is?")
		num := rand.Intn(100) + 1
	guessPoint:
		var guess string
		fmt.Scanln(&guess)
		guessNum, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println("Input is not a valid number. Want to try again?")
			fmt.Println("Options: 1. Try Again 2. Get Answer 3. Quit")
			var resp string
			fmt.Scanln(&resp)
			switch resp {
			case "1":
				fmt.Println("So What number am I thinking about? Will you guess right this time?")
				goto guessPoint
			case "2":
				fmt.Printf("Number is %d. Want to start all over again? 1. Yes,  2. Quit", num)
				var resp2 string
				fmt.Scanln(&resp2)
				if resp2 == "1" {
					continue guessStart
				} else {
					return
				}
			case "3":
				return
			}
		}
		switch {
		case guessNum < num:
			fmt.Printf("Your number is %d less than mine. Want to try again? 1. Yes,  2. Quit\n", num-guessNum)
			var resp2 string
			fmt.Scanln(&resp2)
			switch resp2 {
			case "1":
				fmt.Println("So What number am I thinking about? Will you guess right this time?")
				goto guessPoint
			case "2":
				return
			default:
				goto guessPoint
			}
		case guessNum > num:
			fmt.Printf("Your number is %d more than mine. Want to try again? 1. Yes,  2. Quit\n", guessNum-num)
			var resp2 string
			fmt.Scanln(&resp2)
			switch resp2 {
			case "1":
				fmt.Println("So What number am I thinking about? Will you guess right this time?")
				goto guessPoint
			case "2":
				return
			default:
				goto guessPoint
			}
		case guessNum == num:
			fmt.Println("Correct! You got me. Want to try another number? 1. Yes,  2. Quit")
			var resp2 string
			fmt.Scanln(&resp2)
			if resp2 == "1" {
				continue guessStart
			} else {
				return
			}
		}

	}
}
