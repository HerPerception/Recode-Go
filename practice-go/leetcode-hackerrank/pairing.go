package main

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	count := int32(0)
	for i := range ar {
		for j := i + 1; j < int(n); j++ {
			if ar[i] == 0 || ar[j] == 0 {
				continue
			}
			if ar[i] == ar[j] {
				ar[i] = 0
				ar[j] = 0
				count++
			}

		}
	}
	return count
}

//func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	n := int32(nTemp)

// 	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

// 	var ar []int32

// 	for i := 0; i < int(n); i++ {
// 		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
// 		checkError(err)
// 		arItem := int32(arItemTemp)
// 		ar = append(ar, arItem)
// 	}

// 	result := sockMerchant(n, ar)

// 	fmt.Fprintf(writer, "%d\n", result)

// 	writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

//	func checkError(err error) {
//		if err != nil {
//			panic(err)
//		}
//	}
//fmt.Println(sockMerchant(6, []int32{2, 3, 2, 2, 4, 5}))
//}
