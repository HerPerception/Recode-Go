package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	phonebook := map[string]int{}
	lookup := map[string]bool{}
	scanner := bufio.NewScanner(os.Stdin)
	// count := 0
	// phoneNumbers := make(map[string]int)
	// slice
	for scanner.Scan() {
		str := scanner.Text()
		if len(phonebook) == 0 && len(strings.Fields(str)) == 1 {
			continue
		}
		parts := strings.Fields(str)
		if len(parts) == 2 {
			name := parts[0]
			phone_num, _ := strconv.Atoi(parts[len(parts)-1])
			lookup[name] = true
			phonebook[name] = phone_num
		} else if len(parts) == 1 {
			if !lookup[str] {
				fmt.Println("Not found")
			}
			if lookup[str] {
				fmt.Println(str + "=" + fmt.Sprint(phonebook[str]))
			}
		}
	}
}
