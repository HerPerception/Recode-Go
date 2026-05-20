package main

import "fmt"

func ViewTasks(TaskSlice []string) {
	if len(TaskSlice) == 0 {
		fmt.Println("Ooppss... No tasks added yet. Add one now, would you?")
		return
	} else {
		for i, each := range TaskSlice {
			fmt.Printf("\033[45;30m %d. %s\033[0m\n", i+1, each)
		}
		fmt.Println()
	}
}
