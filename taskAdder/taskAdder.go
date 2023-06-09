package taskAdder

import (
	"fmt"
	"net/http"

	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
	"github.com/Basileus1990/CrowdCompute.git/database"
)

// AddTask adds a task to the database(temporarily in memory)
// TODO: validate the task
// TODO: validate the given code
// TODO: add the task to the database
func AddTask(w http.ResponseWriter, r *http.Request) {
	// print the task name
	newTask := dataStructures.Task{
		Name:        r.FormValue("taskName"),
		Description: r.FormValue("taskDescription"),
		Author:      r.FormValue("taskAuthor"),
		Data:        r.FormValue("taskData"),
		Code:        r.FormValue("taskCode"),
	}

	// Task verification
	if !newTask.VerifyTask() {
		// TODO: Redirect to an error page
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid task")
		return
	}

	database.AddTask(newTask)
	for _, task := range database.GetAllTasks() {
		fmt.Fprintf(w, "%v\n", task)
	}
}
