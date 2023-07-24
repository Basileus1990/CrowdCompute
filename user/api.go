// File contains API endpoints for user management:
//   - login
//   - register
//   - renew token

package user

import (
	"net/http"

	"github.com/Basileus1990/CrowdCompute.git/database"
	"github.com/Basileus1990/CrowdCompute.git/user/user_token"
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
		http.Error(w, ErrWrongLoginOrPass.Error(), http.StatusUnauthorized)
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

func RenewUserToken(w http.ResponseWriter, r *http.Request) {
	token, err := GetTokenFromCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renew the token
	newTokenStr, err := user_token.GenerateToken(token.User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the token to the cookie
	SetUserCookie(w, newTokenStr)

	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Takes in: login, password, email
// Checks if the user already exists
// If not, inserts the user into the database
// func Register(w http.ResponseWriter, r *http.Request) {
// 	// Get the login and password from the request
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
// 	email := r.FormValue("email")
// 	if username == "" || password == "" || email == "" {
// 		http.Error(w, "not all required fields were given", http.StatusBadRequest)
// 		return
// 	}

// 	// Check if the user already exists
// 	_, err := database.GetUserByLogin(username)
// 	if err == nil {
// 		http.Error(w, ErrUserAlreadyExists.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Hash the password
// 	hashedPass, err := hashPassword(password)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Insert the user into the database
// 	err = database.InsertUser(login, hashedPass)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Redirect to the login page
// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
// }
