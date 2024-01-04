package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"todo"
)

func main() {
	var (
		listFlag      bool
		addFlag       string
		markFlag      int
		deleteFlag    int
		deleteAllFlag bool
	)

	flag.BoolVar(&listFlag, "list", false, "List tasks")
	flag.StringVar(&addFlag, "add", "", "Add a new task")
	flag.IntVar(&markFlag, "mark", 0, "Mark a task as done by index")
	flag.IntVar(&deleteFlag, "delete", 0, "Delete a task by index")
	flag.BoolVar(&deleteAllFlag, "deleteall", false, "Delete all tasks")

	// Parse command-line arguments
	flag.Parse()

	todoList := todo.TodoList{}
	filename := "todolist.txt"

	// Load the to-do list from a file (if it exists)
	if err := todoList.LoadFromFile(filename); err != nil {
		fmt.Println("Error loading to-do list:", err)
		os.Exit(1)
	}

	if listFlag {
		// If the -list flag is provided, list tasks and exit
		todoList.ListTasks()
		os.Exit(0)
	}

	if addFlag != "" {
		// If the -add flag is provided, add a new task
		todoList.AddTask(addFlag)
		fmt.Println("Task added successfully.")
	}

	if markFlag > 0 {
		// If the -mark flag is provided, mark a task as done
		if markFlag <= len(todoList.Tasks) {
			task := todoList.Tasks[markFlag-1]
			if !strings.Contains(task, "[x]") {
				todoList.Tasks[markFlag-1] = strings.Replace(task, "[ ]", "[x]", 1)
				fmt.Println("Task marked as done.")
			} else {
				fmt.Println("Task is already marked as done.")
			}
		} else {
			fmt.Println("Invalid task index.")
		}
	}

	if deleteFlag > 0 {
		// If the -delete flag is provided, delete a task
		if deleteFlag <= len(todoList.Tasks) {
			todoList.Tasks = append(todoList.Tasks[:deleteFlag-1], todoList.Tasks[deleteFlag:]...)
			fmt.Println("Task deleted successfully.")
		} else {
			fmt.Println("Invalid task index.")
		}
	}

	if deleteAllFlag {
		// If the -deleteall flag is provided, delete all tasks
		todoList.DeleteAllTasks()
		fmt.Println("All tasks deleted successfully.")
	}

	// Save the updated to-do list to a file
	if err := todoList.SaveToFile(filename); err != nil {
		fmt.Println("Error saving to-do list:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
