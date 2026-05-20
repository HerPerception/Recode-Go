package main

import (
	"strings"
)

/*
 * Complete the 'minimumNumber' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING password
 */

func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"

	num := int32(4)
	if strings.ContainsAny(password, numbers) {
		num -= 1
	}
	if strings.ContainsAny(password, lower_case) {
		num -= 1
	}
	if strings.ContainsAny(password, upper_case) {
		num -= 1
	}
	if strings.ContainsAny(password, special_characters) {
		num -= 1
	}
	sum := n + num
	if sum < 6 {
		less := 6 - sum
		return num + less
	}

	return num
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	n := int32(nTemp)

// 	password := readLine(reader)

// 	answer := minimumNumber(n, password)

// 	fmt.Fprintf(writer, "%d\n", answer)

// 	writer.Flush()
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
