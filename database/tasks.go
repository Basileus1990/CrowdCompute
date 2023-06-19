package database

import (
	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
)

func AddTask(task dataStructures.TaskInfo) error {
	sqlTaskInput := `INSERT INTO tasks
					("title", "author_id", "description", "code")
					VALUES ($1, $2, $3, $4);`
	// gets author_id from DB
	userID, err := GetUserByUsername(task.Author)
	if err != nil {
		return err
	}

	_, err = db.Exec(sqlTaskInput, task.Title, userID, task.Description, task.Code)
	if err != nil {
		return err
	}

	return nil
}

func GetTaskByTitle(title string) dataStructures.TaskInfo {
	sqlGetTask := `SELECT * FROM tasks WHERE title = $1;`
	var task dataStructures.TaskInfo
	err := db.QueryRow(sqlGetTask, title).Scan(&task.Title, &task.Author, &task.Description, &task.Code)
	if err != nil {
		return dataStructures.TaskInfo{}
	}

	return task
}

// Results basic data about all tasks which are available for execution
func GetAllAvailableTasksInfo() ([]dataStructures.TaskInfo, error) {
	sqlGetAllTasks := `SELECT title, author_id, description, code FROM tasks WHERE available = TRUE;`
	rows, err := db.Query(sqlGetAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []dataStructures.TaskInfo
	for rows.Next() {
		author_id := -1
		var task dataStructures.TaskInfo
		err := rows.Scan(&task.Title, &author_id, &task.Description, &task.Code)
		if err != nil {
			return nil, err
		}

		// setting author username
		author, err := GetUserByID(author_id)
		if err != nil {
			return nil, err
		}
		task.Author = author.Username

		tasks = append(tasks, task)
	}
	return tasks, nil
}
