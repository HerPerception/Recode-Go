package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	line := ""
	count := 0
	for scanner.Scan() {
		line = scanner.Text()
		count++
		if count == 1 {
			continue
		}
		res1 := ""
		res2 := ""
		for i, ch := range line {
			if i%2 == 0 {
				res1 += string(ch)
			} else if i%2 != 0 {
				res2 += string(ch)
			}
		}
		fmt.Printf("%s %s\n", res1, res2)
	}

}
