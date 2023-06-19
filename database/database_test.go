package database

import (
	"log"
	"testing"

	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
)

func TestDBInit(t *testing.T) {
	err := db.Ping()
	if err != nil {
		t.Fatal("Database connection failed: ", err)
	}
}

// user struct
type testUser struct {
	username string
	password string
	email    string
}

func deleteTestUsers(users []testUser) {
	sqlDelete := `DELETE FROM users WHERE username = $1;`
	for _, u := range users {
		_, err := db.Exec(sqlDelete, u.username)
		if err != nil {
			log.Println("Error deleting user: ", u.username, err)
		}
	}
}

func deleteTestTasks(tasks []dataStructures.TaskInfo) {
	sqlDelete := `DELETE FROM task_info WHERE title = $1;`
	for _, t := range tasks {
		_, err := db.Exec(sqlDelete, t.Title)
		if err != nil {
			log.Println("Error deleting task: ", t.Title, err)
		}
	}
}

func addTestUsers(users []testUser) {
	sqlUserInput := `INSERT INTO "users" 
					("username", "password", "email") 
					VALUES ($1, $2, $3);`
	for _, u := range users {
		_, err := db.Exec(sqlUserInput, u.username, u.password, u.email)
		if err != nil {
			log.Println("Error adding user: ", u.username, err)
		}
	}
}

func TestAddUser(t *testing.T) {
	sqlUserInput := `INSERT INTO "users" 
					("username", "password", "email") 
					VALUES ($1, $2, $3);`
	sqlResultCount := `SELECT count(*) FROM users;`

	var startRowNumber int
	db.QueryRow(sqlResultCount).Scan(&startRowNumber)

	// test proper adding
	goodUsers := []testUser{
		{"fghfhdfgfhg", "12345678", "fhnfghngfgfg@gmail.com"},
		{"rjynfgnfbfsbfbfs", "12345678", "ujyhdhjufkdfghndgn@gmail.com"},
		{"fgjnfjngnvfgnvdnbfb", "12345678", "ghghmcgnbcgncv@gmail.com"},
	}

	for _, u := range goodUsers {
		_, err := db.Exec(sqlUserInput, u.username, u.password, u.email)
		if err != nil {
			t.Fatal("Database insertion failed: ", sqlUserInput, err)
		}
	}

	// test invalid user data (repeats)
	badUsers := []testUser{
		{"fghfhdfgfhg", "12345678", "fhnfghngfgfg@gmail.com"},
		{"rjynfgnfbfsdfbfbfs", "12345678", "ujyhdhjufkdfghndgn@gmail.com"},
		{"fgjnfjngnvfgnvdnbfb", "12345678", "ghghmcdfgnbcgncv@gmail.com"},
	}
	for _, u := range badUsers {
		_, err := db.Exec(sqlUserInput, u.username, u.password, u.email)
		if err == nil {
			deleteTestUsers(badUsers)
			t.Fatal("Database insertion should fail: ", sqlUserInput, err)
		}
	}

	deleteTestUsers(goodUsers)

	// test if state is the same
	var endRowNumber int
	db.QueryRow(sqlResultCount).Scan(&endRowNumber)
	if startRowNumber != endRowNumber {
		t.Fatal("Database changed state after the test where it shouldn't. Before num: ", startRowNumber, " After num: ", endRowNumber)
	}
}

func TestAddTasksInfo(t *testing.T) {
	sqlTaskInput := `INSERT INTO "task_info"
					("title", "author_id", "description", "code")
					VALUES ($1, $2, $3, $4);`

	goodUsers := []testUser{
		{"fghfhdfgfhg", "12345678", "fhnfghngfgfg@gmail.com"},
		{"rjynfgnfbfsbfbfs", "12345678", "ujyhdhjufkdfghndgn@gmail.com"},
		{"fgjnfjngnvfgnvdnbfb", "12345678", "ghghmcgnbcgncv@gmail.com"},
	}
	addTestUsers(goodUsers)
	defer deleteTestUsers(goodUsers)

	goodTask := []dataStructures.TaskInfo{
		{Title: "title1", Author: "fghfhdfgfhg", Description: "description1", Code: "code1"},
		{Title: "title2", Author: "rjynfgnfbfsbfbfs", Description: "description2", Code: "code2"},
		{Title: "title3", Author: "fgjnfjngnvfgnvdnbfb", Description: "description3", Code: "code3"},
	}
	sampleUser, err := GetUserByUsername(goodUsers[0].username)
	if err != nil {
		t.Fatal("Error getting user: ", err)
	}
	for _, task := range goodTask {
		_, err := db.Exec(sqlTaskInput, task.Title, sampleUser.ID, task.Description, task.Code)
		if err != nil {
			t.Fatal("Database insertion failed: ", sqlTaskInput, err)
		}
	}
	defer deleteTestTasks(goodTask)

	// test invalid user data
	badTasks := []dataStructures.TaskInfo{
		{Title: "title1", Author: "fghfhdfgfhg", Description: "description", Code: "code"},
		{Title: "title2", Author: "rjynfgnfbfsbfbfs", Description: "description", Code: "code"},
	}
	for _, task := range badTasks {
		_, err := db.Exec(sqlTaskInput, task.Title, sampleUser.ID, task.Description, task.Code)
		if err == nil {
			deleteTestTasks(badTasks)
			t.Fatal("Database insertion should fail: ", sqlTaskInput, err)
		}
	}

}
