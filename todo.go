package todo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	emptyTask     = "[ ]"
	completedTask = "[x]"
)

// TodoList struct represents the to-do list
type TodoList struct {
	Tasks []string
}

// AddTask adds a new task to the to-do list
func (t *TodoList) AddTask(description string) {
	task := fmt.Sprintf("%s %s", emptyTask, description)
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
func (t *TodoList) MarkOrDeleteTask(input string) error {
	parts := strings.Fields(input)
	if len(parts) < 2 {
		return errors.New("Invalid input. Please provide a valid task index.")
	}

	index, err := strconv.Atoi(parts[1])
	if err != nil {
		return errors.New("Invalid input. Please provide a valid task index.")
	}

	command := strings.ToLower(parts[0])
	switch command {
	case "mark":
		return t.markTask(index)
	case "delete":
		return t.deleteTask(index)
	default:
		return errors.New("Invalid command. Please use 'mark' or 'delete' followed by the task index.")
	}
}

func (t *TodoList) markTask(index int) error {
	if index < 1 || index > len(t.Tasks) {
		return errors.New("Invalid task index.")
	}

	task := t.Tasks[index-1]
	if !strings.Contains(task, completedTask) {
		t.Tasks[index-1] = strings.Replace(task, emptyTask, completedTask, 1)
		fmt.Println("Task marked as done.")
	} else {
		fmt.Println("Task is already marked as done.")
	}
	return nil
}

// deleteTask deletes a specific task from the to-do list
func (t *TodoList) deleteTask(index int) error {
	if index < 1 || index > len(t.Tasks) {
		return errors.New("Invalid task index.")
	}

	t.Tasks = append(t.Tasks[:index-1], t.Tasks[index:]...)
	fmt.Println("Task deleted successfully.")
	return nil
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
