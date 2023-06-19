package database

import (
	"log"
	"testing"
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

func TestAddUser(t *testing.T) {
	sqlUserInput := `INSERT INTO "users" 
					("username", "password", "email") 
					VALUES ($1, $2, $3);`
	sqlResultCount := `SELECT count(*) FROM users;`

	var startRowNumber int
	db.QueryRow(sqlResultCount).Scan(&startRowNumber)
	log.Println(startRowNumber)

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

// func TestAddTasks(t *testing.T) {
// 	sqlUserInput := `INSERT INTO "users"
// 					("username", "password", "email")
// 					VALUES ($1, $2, $3);`
// 	sqlTaskInput := `INSERT INTO "users"
// 					("title", "author_id", "description", "code")
// 					VALUES ($1, $2, $3, $4);`

// 	goodUsers := []testUser{
// 		{"fghfhdfgfhg", "12345678", "fhnfghngfgfg@gmail.com"},
// 		{"rjynfgnfbfsbfbfs", "12345678", "ujyhdhjufkdfghndgn@gmail.com"},
// 		{"fgjnfjngnvfgnvdnbfb", "12345678", "ghghmcgnbcgncv@gmail.com"},
// 	}
// 	for _, u := range goodUsers {
// 		_, err := db.Exec(sqlUserInput, u.username, u.password, u.email)
// 		if err != nil {
// 			t.Fatal("Database insertion failed: ", sqlUserInput, err)
// 		}
// 	}
// 	defer deleteTestUsers(goodUsers)

// 	goodTask := []dataStructures.Task{
// 		{"title1", "fghfhdfgfhg", "description1", "code1"},
// 		{"title2", "rjynfgnfbfsbfbfs", "description2", "code2"},
// 		{"title3", "fgjnfjngnvfgnvdnbfb", "description3", "code3"},
// 	}
// 	// change DB structure
// }
