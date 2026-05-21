package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var TaskSlice []string
	var NewTasks []string
TaskLoop:
	for {
		fmt.Println("options: 1. Add Task, 2. View Tasks, 3. Deleted Completed tasks, 4. Exit")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option := scanner.Text()
		if option == "4" {
			return
		}
		if len(option) != 1 {
			continue TaskLoop
		} else if option != "1" && option != "2" && option != "3" {
			continue TaskLoop
		}

		if option == "1" {
			NewTasks = AddTask(TaskSlice)
			continue TaskLoop
		}
		if option == "2" {
			ViewTasks(NewTasks)
		}
		if option == "3" {
			var err error
			NewTasks, err = DeleteTask(NewTasks)
			if err != nil {
				fmt.Println(err)
				ViewTasks(NewTasks)
			}
			continue TaskLoop
		}
	}
}
