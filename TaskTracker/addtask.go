package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddTask(TaskSlice []string) []string {
	fmt.Println("Write Task Below, Write 'Done' to add task(s):")
	scanner := bufio.NewScanner(os.Stdin)
	var userText string
	for scanner.Scan() {
		userText = scanner.Text()
		if strings.ToLower(userText) == "done" {
			break
		} else {
			TaskSlice = append(TaskSlice, userText)
		}
	}
	if scanner.Err() != nil {
		fmt.Println("There was an error adding your task. Please try again.")
		return []string{}
	}
	fmt.Println("Task added successfully.")
	return TaskSlice
}
