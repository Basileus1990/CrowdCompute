// File responsible for all database operations related to users.

package database

import (
	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
)

// TODO: REFACTOR THIS

func GetUserByUsername(username string) (dataStructures.User, error) {
	sqlGetUserID := `SELECT email FROM users WHERE username = $1;`
	var user dataStructures.User
	user.Username = username
	err := db.QueryRow(sqlGetUserID, username).Scan(&user.Email)
	if err != nil {
		return dataStructures.User{}, err
	}

	return user, nil
}

func AddUser(user dataStructures.User, pass string) error {
	sqlAddUser := `INSERT INTO users(username, email, password) VALUES($1, $2, $3);`
	_, err := db.Exec(sqlAddUser, user.Username, user.Email, pass)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(username string) error {
	sqlDeleteUser := `DELETE FROM users WHERE username = $1;`
	_, err := db.Exec(sqlDeleteUser, username)
	if err != nil {
		return err
	}
	return nil
}

func UserExists(username string) (bool, error) {
	sqlUserExists := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1);`
	var exists bool
	err := db.QueryRow(sqlUserExists, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func GetPasswordByLogin(login string) (string, error) {
	sqlGetPassword := `SELECT password FROM users WHERE username = $1 OR email = $1;`
	var password string
	err := db.QueryRow(sqlGetPassword, login).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}

func GetUserByLogin(login string) (dataStructures.User, error) {
	sqlGetUser := `SELECT username, email FROM users WHERE username = $1 OR email = $1;`
	var user dataStructures.User
	err := db.QueryRow(sqlGetUser, login).Scan(&user.Username, &user.Email)
	if err != nil {
		return dataStructures.User{}, err
	}

	return user, nil
}
