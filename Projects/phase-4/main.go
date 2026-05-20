package main

import (
	"bufio"
	"fmt"
	"os"

	"sentinel/myDir"
)

func main() {
	fmt.Println("SENTINEL - COMMAND & CONTROL CONSOLE")
hackLoop:
	for {
		fmt.Println("All systems nominal. str -> string transformer, calc -> calculator, base -> base, history -> last 10 results, clear -. clear previous sessions, exit -> close session and console.")
		fmt.Println("Type 'help' to see commands and functions if you are unsure where to begin.")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if input == "help" {
			fmt.Println("calc → the calculator")
			fmt.Println("base → the base converter")
			fmt.Println("str → the string transformer")
			fmt.Println("help → shows all commands")
			fmt.Println("history → shows last 10 inputs")
			fmt.Println("exit → shuts down the session")
			continue hackLoop
		}
		for {
			if input == "calc" {
				fmt.Println("Choose options: add <a b>, sub <a b>, mul <a b>, div <a b>, mod <a b>, pow <a b>, last -> shows last result, back -> return.")
				scanner.Scan()
				newInput := scanner.Text()
				if newInput == "back" {
					myDir.Back(newInput)
					continue hackLoop
				} else {
					fmt.Println(myDir.Calculator(newInput))
				}
			}
			if input == "base" {
				fmt.Println("Choose options: hex <a>, bin <a>, dec <a>, last -> shows last result, back -> return.")
				scanner.Scan()
				newInput := scanner.Text()
				if newInput == "back" {
					myDir.Back(newInput)
					continue hackLoop
				} else {
					fmt.Println(myDir.Base(newInput))
				}
			}
			if input == "str" {
				fmt.Println("Choose options: upper <text>, low <text>, cap <text>, title <text>, snake <text>, reverse <text>, last -> shows last result, back -> return.")
				scanner.Scan()
				newInput := scanner.Text()
				if newInput == "back" {
					myDir.Back(newInput)
					continue hackLoop
				}
				fmt.Println(myDir.StringsTran(newInput))

			}

			if input == "history" {
				fmt.Println(myDir.PrintHistory(input))
				continue hackLoop
			}
			if input == "clear" {
				fmt.Println("Clearing previous session....")
				fmt.Println(myDir.PrintHistory(input))
				continue hackLoop
			}
			if input == "exit" {
				fmt.Println("Shutting down console. Goodbye...")
				return
			}
			if input == "back" {
				continue hackLoop
			}
		}
	}
}
