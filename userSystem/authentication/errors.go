package authentication

import "fmt"

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
