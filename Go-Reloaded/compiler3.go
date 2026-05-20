package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    result := "I have 42 (hex) numbers of files and 1010 (bin) folders on my desktop."

    bases := map[string]int{
        "(hex)": 16,
        "(bin)": 2,
    }

    words := strings.Fields(result)

    for i := 0; i < len(words); i++ {
        if base, ok := bases[words[i]]; ok {
            num, err := strconv.ParseInt(words[i-1], base, 64)
            if err != nil {
                fmt.Println("Error parsing number:", err)
                continue
            }
            words[i-1] = fmt.Sprintf("%d", num)
            words = append(words[:i], words[i+1:]...)
            i-- // adjust index
        }
    }

    final := strings.Join(words, " ")
    fmt.Println(final)
}