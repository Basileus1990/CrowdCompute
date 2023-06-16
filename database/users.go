package database

func GetUserIDByUsername(username string) (int, error) {
	sqlGetUserID := `SELECT id FROM users WHERE username = $1;`
	var userID int
	err := db.QueryRow(sqlGetUserID, username).Scan(&userID)
	if err != nil {
		return -1, err
	}
	return userID, nil
}

func GetUsernameByID(id int) (string, error) {
	sqlGetUsername := `SELECT username FROM users WHERE id = $1;`
	var username string
	err := db.QueryRow(sqlGetUsername, id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
