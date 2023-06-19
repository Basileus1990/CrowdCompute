package dataStructures

import "encoding/json"

// Task is a struct that represents a task
// Requrirements:
//   - Name && Author && Code != ""
//   - Data - a valid JSON string
type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Code        string `json:"code"`
	Data        string `json:"data"`
}

func (task Task) VerifyTask() bool {
	if task.Name == "" || task.Author == "" || task.Code == "" {
		return false
	}
	if task.Data != "" && !json.Valid([]byte(task.Data)) {
		return false
	}
	return true
}
