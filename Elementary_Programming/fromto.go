package main

import (
	"fmt"
	"strconv"
)

func FromTo(from int64, to int64) string {
	if (from < 0 || from > 99) || to < 0 || to > 99 {
		return "Invalid\n"
	}
	word := ""
	if from == to {
		return fmt.Sprintf("%s\n", strconv.FormatInt(from, 10))
	}
	if from > to {
		for i := from; i >= to; i-- {
			if i < 10 && i != to {
				word += string('0') + strconv.FormatInt(i, 10) + ", "
			} else if i >= 10 && i != to {
				word += strconv.FormatInt(i, 10) + ", "
			}
			if i == to && i < 10 {
				word += string('0') + strconv.FormatInt(i, 10)
			} else if i == to && i >= 10 {
				word += strconv.FormatInt(i, 10)
			}
		}
	} else {
		for i := from; i <= to; i++ {
			if i < 10 && i != to {
				word += string('0') + strconv.FormatInt(i, 10) + ", "
			} else if i >= 10 && i != to {
				word += strconv.FormatInt(i, 10) + ", "
			}
			if i == to && i < 10 {
				word += string('0') + strconv.FormatInt(i, 10)
			} else if i == to && i >= 10 {
				word += strconv.FormatInt(i, 10)
			}
		}

	}
	s := fmt.Sprintf("%s\n", word)
	return s
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
