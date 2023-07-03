package dataStructures

// TaskInfo is a struct that represents a task basic information
// Requrirements:
//   - Name && Author && Code != ""
type TaskInfo struct {
	Title       string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

func (task TaskInfo) VerifyTask() error {
	if task.Title == "" {
		return &InvalidDataError{Field: "name", Data: task.Title}
	}
	if task.Author == "" {
		return &InvalidDataError{Field: "author", Data: task.Author}
	}
	if task.Code == "" {
		return &InvalidDataError{Field: "code", Data: task.Code}
	}
	return nil
}

// Task is a struct that represents a single task and its data
type Task struct {
	TaskTitle string `json:"task_title"`
	Data      string `json:"data"`
}

// Result is a struct that represents a single task result
// Requirements:
//   - Results is a valid JSON
type Result struct {
	TaskID           int    `json:"task_id"`
	ExecutorUsername string `json:"executor_username"`
	Results          string `json:"results"`
}
