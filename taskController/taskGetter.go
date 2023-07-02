package taskController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
	"github.com/Basileus1990/CrowdCompute.git/database"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := database.GetAllAvailableTasksInfo()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json, err := tasksToJSON(tasks)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(json)
}

func tasksToJSON(tasks []dataStructures.TaskInfo) ([]byte, error) {
	json, err := json.Marshal(tasks)
	if err != nil {
		return nil, err
	}
	return json, nil
}
