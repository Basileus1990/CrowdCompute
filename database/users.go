package database

import (
	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
)

func GetUserByUsername(username string) (dataStructures.User, error) {
	sqlGetUserID := `SELECT id, email FROM users WHERE username = $1;`
	var user dataStructures.User
	user.Username = username
	err := db.QueryRow(sqlGetUserID, username).Scan(&user.ID, &user.Email)
	if err != nil {
		return dataStructures.User{}, err
	}

	return user, nil
}

func GetUserByID(id int) (dataStructures.User, error) {
	sqlGetUsername := `SELECT username, email FROM users WHERE id = $1;`
	var user dataStructures.User
	user.ID = id
	err := db.QueryRow(sqlGetUsername, id).Scan(&user.Username, &user.Email)
	if err != nil {
		return dataStructures.User{}, err
	}

	return user, nil
}
