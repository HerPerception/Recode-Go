package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	//count := int(n - 1)
	//n2 := 1
	for i := 1; i <= int(n); i++ {
		spaces := strings.Repeat(" ", int(n-int32(i)))
		s := strings.Repeat("#", i)
		fmt.Printf("%s%s\n", spaces, s)
		//count -= 1
		//n2 += 1
	}
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	n := int32(nTemp)

// 	staircase(n)
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	staircase(10)
}
