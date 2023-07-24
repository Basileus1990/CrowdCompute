package usersystem

import (
	"github.com/Basileus1990/CrowdCompute.git/database"
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(login, password string) (bool, error) {
	hashPass, err := database.GetPasswordByLogin(login)
	if err != nil {
		return false, err
	}

	hashedNewPass, err := hashPassword(password)
	if err != nil {
		return false, err
	}
	return hashPass == hashedNewPass, nil
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}
