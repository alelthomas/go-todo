package todo

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// TodoList struct represents the to-do list
type TodoList struct {
	Tasks []string
}

// AddTask adds a new task to the to-do list
func (t *TodoList) AddTask(description string) {
	task := fmt.Sprintf("[ ] %s", description)
	t.Tasks = append(t.Tasks, task)
	fmt.Println("Task added successfully.")
}

// ListTasks prints all tasks in the to-do list
func (t *TodoList) ListTasks() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks in the to-do list.")
		return
	}

	fmt.Println("~ My 2024 goals ~")
	for i, task := range t.Tasks {
		parts := strings.Fields(task)
		if len(parts) >= 3 && parts[2] != "" {
			fmt.Printf("%d. %s\n", i+1, task)
		} else {
			fmt.Printf("%d. [ ]\n", i+1)
		}
	}
}

// MarkOrDeleteTask marks a task as done or deletes it based on the input index
func (t *TodoList) MarkOrDeleteTask(input string) {
	parts := strings.Fields(input)
	if len(parts) < 2 {
		fmt.Println("Invalid input. Please provide a valid task index.")
		return
	}

	index, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid task index.")
		return
	}

	command := strings.ToLower(parts[0])
	switch command {
	case "mark":
		if index >= 1 && index <= len(t.Tasks) {
			task := t.Tasks[index-1]
			if !strings.Contains(task, "[x]") {
				t.Tasks[index-1] = strings.Replace(task, "[ ]", "[x]", 1)
				fmt.Println("Task marked as done.")
			} else {
				fmt.Println("Task is already marked as done.")
			}
		} else {
			fmt.Println("Invalid task index.")
		}
	case "delete":
		if index >= 1 && index <= len(t.Tasks) {
			t.Tasks = append(t.Tasks[:index-1], t.Tasks[index:]...)
			fmt.Println("Task deleted successfully.")
		} else {
			fmt.Println("Invalid task index.")
		}
	default:
		fmt.Println("Invalid command. Please use 'mark' or 'delete' followed by the task index.")
	}
}

// DeleteAllTasks deletes all tasks from the to-do list
func (t *TodoList) DeleteAllTasks() {
	t.Tasks = nil
	fmt.Println("All tasks deleted successfully.")
}

// SaveToFile saves the to-do list to a file
func (t *TodoList) SaveToFile(filename string) error {
	data := []byte(strings.Join(t.Tasks, "\n"))
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}

// LoadFromFile loads the to-do list from a file
func (t *TodoList) LoadFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	t.Tasks = strings.Split(string(data), "\n")
	return nil
}
