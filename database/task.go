// File responsible for all database operations related to tasks.

package database

import (
	"encoding/json"
	"errors"

	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
)

// Adds task's information describing the task to the DB. Requred before adding any tasks.
func AddTaskInfo(task dataStructures.TaskInfo) error {
	sqlTaskInput := `INSERT INTO task_info
					("title", "author_username", "description", "code")
					VALUES ($1, $2, $3, $4);`

	_, err := db.Exec(sqlTaskInput, task.Title, task.Author, task.Description, task.Code)
	if err != nil {
		return err
	}

	return nil
}

func AddTask(task dataStructures.Task) error {
	sqlTaskInput := `INSERT INTO task
					("task_info_title", "data")
					VALUES ($1, $2);`

	if !json.Valid([]byte(task.Data)) {
		return errors.New("invalid data in a task")
	}

	_, err := db.Exec(sqlTaskInput, task.TaskTitle, task.Data)
	if err != nil {
		return err
	}

	return nil
}

// Gets task's information describing the task from the DB searching by the title.
func GetTaskInfoByTitle(title string) (dataStructures.TaskInfo, error) {
	sqlGetTask := `SELECT title, author_username, description, code FROM task_info WHERE title = $1;`
	var task dataStructures.TaskInfo
	err := db.QueryRow(sqlGetTask, title).Scan(&task.Title, &task.Author, &task.Description, &task.Code)
	if err != nil {
		return dataStructures.TaskInfo{}, err
	}

	return task, nil
}

// Results basic data about all tasks which are available for execution
func GetAllAvailableTasksInfo() ([]dataStructures.TaskInfo, error) {
	sqlGetAllTasks := `SELECT title, author_username, description, code FROM task_info;`
	rows, err := db.Query(sqlGetAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []dataStructures.TaskInfo
	for rows.Next() {
		var task dataStructures.TaskInfo
		err := rows.Scan(&task.Title, &task.Author, &task.Description, &task.Code)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}
