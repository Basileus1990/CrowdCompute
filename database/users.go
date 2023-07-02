// File responsible for all database operations related to users.

package database

import (
	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
)

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
