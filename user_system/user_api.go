package usersystem

import (
	"net/http"

	"github.com/Basileus1990/CrowdCompute.git/database"
	"github.com/Basileus1990/CrowdCompute.git/user_system/user_token"
)

// Takes the login and password from the request and checks if they are correct.
// If they are correct generates a new token and sets it to the cookie.
func Login(w http.ResponseWriter, r *http.Request) {
	// Get the login and password from the request
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login == "" || password == "" {
		http.Error(w, "login or password not provided", http.StatusBadRequest)
		return
	}

	// Check if the login and password are correct
	ok, err := CheckPassword(login, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		// If they are not correct return an error
		http.Error(w, WrongLoginOrPassError.Error(), http.StatusUnauthorized)
		return
	}

	// Generate a new token
	user, err := database.GetUserByLogin(login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := user_token.GenerateToken(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the token to the cookie
	SetUserCookie(w, token)

	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
