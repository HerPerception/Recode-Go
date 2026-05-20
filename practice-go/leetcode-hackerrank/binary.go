package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	res := []int32{}
	for n > 0 {
		digit := n % 2
		res = append(res, digit)
		n = n / 2
	}
	//  fmt.Println(res)
	count := 0
	count2 := 0
	for i := range res {
		if res[i] == 0 {
			if count2 < count {
				count2 = count
				count = 0
			} else {
				count = 0
			}
		} else if res[i] == 1 {
			count++
		}
	}
	if count2 > count {
		fmt.Println(count2)
	} else {
		count2 = count
		fmt.Println(count2)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
