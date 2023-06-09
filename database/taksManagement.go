package database

import (
	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
)

func AddTask(task dataStructures.Task) {
	db[task.Name] = task
}

func GetTask(name string) dataStructures.Task {
	return db[name]
}

func GetAllTasks() []dataStructures.Task {
	var tasks []dataStructures.Task
	for _, task := range db {
		tasks = append(tasks, task)
	}
	return tasks
}
