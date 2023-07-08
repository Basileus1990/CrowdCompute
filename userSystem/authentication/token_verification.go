// Problems:
// 1. Token can contain a "." which will break the token
// 2. Problem with wrong token being saved in db or something like this
// 3. Make it more DRY
//     * spliting the token

package authentication

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Basileus1990/CrowdCompute.git/database"
)

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

// Returns the username if the token is valid
// Returns an error if the token is invalid or expired
func VerifyToken(token string) (string, error) {
	// decrypt the token
	decryptedToken, err := decryptToken(token)
	if err != nil {
		return "", err
	}
	log.Println("Decrypted: ", decryptedToken)

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
	splittedToken := strings.Split(decryptedToken, "\"}.")
	splittedToken[0] = splittedToken[0] + "\"}"
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

	// check if the token is asigned to the user
	if err := checkUserAsignedToken(tokenStruct.Username, decryptedToken); err != nil {
		return err
	}

	return nil
}

// Returns an error if provided the token is not asigned to the user
func checkUserAsignedToken(username string, token string) error {
	// get the token from the database
	dbToken, err := database.GetAuthToken(username)
	if err != nil {
		return err
	}
	dbToken, err = decryptToken(dbToken)
	if err != nil {
		return err
	}

	log.Println(string(dbToken))
	log.Println(token + "\n")

	// check if the token is the same
	if string(dbToken) != token {
		return &ErrInvalidToken{Msg: "Asigned token is not the same as the one provided"}
	}
	return nil
}

func checkSignature(token string, signature string) error {
	sig, err := getSignature([]byte(token))
	if err != nil {
		return err
	}

	if string(sig) != signature {
		return &ErrInvalidToken{Msg: "Invalid signature, token has been tampered with"}
	}
	return nil
}

// check if the token has expired
func checkExpiration(tokenExp string) error {
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
	splittedToken := strings.Split(token, "\"}.")
	splittedToken[0] = splittedToken[0] + "\"}"
	var tokenStruct Token
	err := json.Unmarshal([]byte(splittedToken[0]), &tokenStruct)
	if err != nil {
		return "", err
	}
	return tokenStruct.Username, nil
}
