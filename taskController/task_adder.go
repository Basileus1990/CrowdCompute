package taskController

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
	newTask := dataStructures.TaskInfo{
		Title:       r.FormValue("taskName"),
		Description: r.FormValue("taskDescription"),
		Author:      r.FormValue("taskAuthor"),
		Code:        r.FormValue("taskCode"),
	}

	// Task verification
	err := newTask.VerifyTask()
	if err != nil {
		// TODO: Redirect to an error page
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	err = database.AddTaskInfo(newTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Couldn't add the task: %v\n%s", newTask, err)
		return
	}

	fmt.Fprintf(w, "Task added")
	tasks, err := database.GetAllAvailableTasksInfo()
	if err != nil {
		fmt.Fprintf(w, "Couldn't get all tasks: %v\n", err)
		return
	}
	for _, task := range tasks {
		fmt.Fprintf(w, "%v\n", task)
	}
}
