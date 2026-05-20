package main

import (
    "fmt"
    "strings"
    //"unicode"
    )
    
func main() {
    text := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
    
    newText := strings.Fields(text)
    for i := 0; i < len(newText); i++ {
        if newText[i] == "(cap)" {
            marker := strings.Trim(newText[i], "()")
            newMarker := strings.Split(marker, ",")
            if len(newMarker) > 1 
            newerText := strings.Title(newText[i-1])
            newText[i-1] = fmt.Sprintf("%v", newerText)
            newText = append(newText[:i], newText[i+1:]...)
            i--
        }
    }
    final := strings.Join(newText, " ")
    fmt.Println(final)
}