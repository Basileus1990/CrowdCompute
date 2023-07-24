package user

import "errors"

// UserNotFoundError is returned when the user has not been found in the database
var ErrUserNotFound error = errors.New("user has not been found")

var ErrUserAlreadyExists error = errors.New("user already exists")

var ErrWrongLoginOrPass error = errors.New("wrong login or password")
