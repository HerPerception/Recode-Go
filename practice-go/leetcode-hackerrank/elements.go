package main

import (
    //"bufio"
   // "fmt"
   // "io"
   // "os"
   // "strconv"
   // "strings"
)



/*
 * Complete the 'countResponseTimeRegressions' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY responseTimes as parameter.
 */

func countResponseTimeRegressions(responseTimes []int32) int32 {
    // Write your code here
	sum := int64(0)
	count := int32(0)
	match := int32(0)
	for i, num := range responseTimes {
		div := float32(0)
		if i > 0 {
			div = float32(sum / int64(count))
		}
		if i > 0 && float32(num) > div {
			match += 1
		}
		sum += int64(num)
		count += 1
		
	}
	return match
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     responseTimesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var responseTimes []int32

//     for i := 0; i < int(responseTimesCount); i++ {
//         responseTimesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         responseTimesItem := int32(responseTimesItemTemp)
//         responseTimes = append(responseTimes, responseTimesItem)
//     }

//     result := countResponseTimeRegressions(responseTimes)

//     fmt.Printf("%d\n", result)
// }

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }
