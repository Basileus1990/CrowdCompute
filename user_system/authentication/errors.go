package authentication

import "fmt"

type ErrInvalidTokenData struct {
	Msg string
}

func (e *ErrInvalidTokenData) Error() string {
	if e.Msg == "" {
		return "Invalid authentication token data"
	} else {
		return fmt.Sprintf("Invalid authentication token data: %s", e.Msg)
	}
}

type ErrInvalidToken struct {
	Msg string
}

func (e *ErrInvalidToken) Error() string {
	if e.Msg == "" {
		return "Invalid authentication token"
	} else {
		return fmt.Sprintf("Invalid authentication token: %s", e.Msg)
	}
}
