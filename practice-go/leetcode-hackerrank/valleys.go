package main

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	// 1. U -> Altitude is now 1 (On a mountain)
	// 2. D -> Altitude is now 0 (Back at sea level, finished a mountain)
	// 3. D -> Altitude is now -1 (Entered a valley!)
	// 4. D -> Altitude is now -2 (Deeper in the valley)
	// 5. U -> Altitude is now -1 (Climbing up, still in the valley)
	// 6. D -> Altitude is now -2 (Going down again)
	// 7. U -> Altitude is now -1 (Climbing up)
	// 8. U -> Altitude is now 0 (Returned to sea level! +1 Valley counted)
	// 9. Total Valleys = 1

	valley := int32(0)
	alt := int32(0)
	for _, ch := range path {
		if ch == 'U' {
			alt += 1
			if alt == 0 {
				valley += 1
			}
		}
		if ch == 'D' {
			alt -= 1
		}

	}
	return valley
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//     checkError(err)

//     defer stdout.Close()

//     writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

//     stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)
//     steps := int32(stepsTemp)

//     path := readLine(reader)

//     result := countingValleys(steps, path)

//     fmt.Fprintf(writer, "%d\n", result)

//     writer.Flush()
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
