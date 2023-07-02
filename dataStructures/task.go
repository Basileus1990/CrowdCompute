package dataStructures

// TaskInfo is a struct that represents a task basic information
// Requrirements:
//   - Name && Author && Code != ""
//   - Data - a valid JSON string
type TaskInfo struct {
	Title       string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

type Task struct {
	TaskTitle string `json:"task_title"`
	Data      string `json:"data"`
}

func (task TaskInfo) VerifyTask() bool {
	if task.Title == "" || task.Author == "" || task.Code == "" {
		return false
	}
	return true
}
