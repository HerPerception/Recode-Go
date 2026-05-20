package main

import "fmt"

func DeleteTask(TaskSlice []string) []string {
	fmt.Println("Type ID (Number) of task you want to delete.")
	ID := 0
	fmt.Scanln(&ID)
	if ID < 0 || ID > len(TaskSlice) {
		fmt.Println("Oooopppsss... ID does not exist. Below are the tasks and their IDs:")
		return TaskSlice
	}
	TaskSlice = append(TaskSlice[:ID-1], TaskSlice[ID:]...)
	fmt.Println("Task deleted successfully.")
	return TaskSlice
}
