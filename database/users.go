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

func GetAuthToken(username string) (string, error) {
	sqlGetAuthToken := `SELECT auth_token FROM users WHERE username = $1;`
	var token string
	err := db.QueryRow(sqlGetAuthToken, username).Scan(&token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func SetAuthToken(username string, token string) error {
	sqlSetAuthToken := `UPDATE users SET auth_token = $1 WHERE username = $2;`
	_, err := db.Exec(sqlSetAuthToken, token, username)
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
