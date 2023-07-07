package authentication

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ErrInvalidToken struct {
	Msg string
}

func (e *ErrInvalidToken) Error() string {
	if e.Msg != "" {
		return "Invalid authentication token"
	} else {
		return fmt.Sprintf("Invalid authentication token: %s", e.Msg)
	}
}

// Returns the username if the token is valid
// Returns an error if the token is invalid or expired
func VerifyToken(token string) (string, error) {
	// decrypt the token
	decryptedToken, err := decryptToken(token)
	if err != nil {
		return "", err
	}

	// check if the token is valid
	if err = validateToken(decryptedToken); err != nil {
		return "", err
	}

	// return the username
	username, err := getUsername(decryptedToken)
	if err != nil {
		return "", err
	}
	return username, nil
}

func validateToken(decryptedToken string) error {
	splittedToken := strings.Split(decryptedToken, ".")
	if len(splittedToken) != 2 {
		return &ErrInvalidToken{}
	}

	// check if the signature is valid
	if err := checkSignature(splittedToken[0], splittedToken[1]); err != nil {
		return err
	}

	// create an util struct to get the expiration time
	var tokenStruct Token
	err := json.Unmarshal([]byte(splittedToken[0]), &tokenStruct)
	if err != nil {
		return err
	}

	// check if the token is expired
	if err := checkExpiration(tokenStruct.ExpirationTimeStamp); err != nil {
		return err
	}
	return nil
}

func checkSignature(token string, signature string) error {
	// get the signature
	sig, err := getSignature([]byte(token))
	if err != nil {
		return err
	}

	if string(sig) != signature {
		return &ErrInvalidToken{Msg: "Invalid signature, token has been tampered with"}
	}
	return nil
}

func checkExpiration(tokenExp string) error {
	// check if the token has expired
	expTime, err := time.ParseDuration(fmt.Sprintf("%sns", tokenExp))
	if err != nil {
		return err
	}
	if expTime.Nanoseconds() < time.Now().UnixNano() {
		return &ErrInvalidToken{Msg: "Token has expired"}
	}
	return nil
}

func getUsername(token string) (string, error) {
	var tokenStruct Token
	err := json.Unmarshal([]byte(token), &tokenStruct)
	if err != nil {
		return "", err
	}
	return tokenStruct.Username, nil
}
