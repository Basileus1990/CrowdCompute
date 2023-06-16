package database

import (
	"log"

	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
)

func AddTask(task dataStructures.Task) error {
	sqlTaskInput := `INSERT INTO tasks
					("title", "author_id", "description", "code")
					VALUES ($1, $2, $3, $4);`
	// gets author_id from DB
	userID, err := GetUserIDByUsername(task.Author)
	if err != nil {
		return err
	}

	_, err = db.Exec(sqlTaskInput, task.Name, userID, task.Description, task.Code)
	if err != nil {
		return err
	}

	return nil
}

func GetTask(title string) dataStructures.Task {
	sqlGetTask := `SELECT * FROM tasks WHERE title = $1;`
	var task dataStructures.Task
	err := db.QueryRow(sqlGetTask, title).Scan(&task.Name, &task.Author, &task.Description, &task.Code)
	if err != nil {
		return dataStructures.Task{}
	}
	return task
}

func GetAllTasks() ([]dataStructures.Task, error) {
	sqlGetAllTasks := `SELECT title, author_id, description, code FROM tasks;`
	rows, err := db.Query(sqlGetAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []dataStructures.Task
	log.Println(tasks)
	for rows.Next() {
		author_id := -1
		var task dataStructures.Task
		err := rows.Scan(&task.Name, &author_id, &task.Description, &task.Code)
		if err != nil {
			return nil, err
		}
		task.Author, err = GetUsernameByID(author_id)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
