package usersystem

import "errors"

// UserNotFoundError is returned when the user has not been found in the database
var UserNotFoundError error = errors.New("User has not been found")

var UserAlreadyExistsError error = errors.New("User already exists")

var WrongLoginOrPassError error = errors.New("Wrong login or password")
