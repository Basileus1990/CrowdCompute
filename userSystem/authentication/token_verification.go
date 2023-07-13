// Problems:
// 1. Token can contain a "." which will break the token
// 2. Problem with wrong token being saved in db or something like this
// 3. Make it more DRY
//     * spliting the token

package authentication

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Basileus1990/CrowdCompute.git/database"
)

// Checks if the given encrypted signed token is valid.
//   - checks if the signature is valid
//   - checks if the signed token has expired
//   - checks if the given signed token is asigned to the user
//
// If valid, it returns the Token struct. Otherwise it returns an error.
func VerifyToken(encToken string) (*Token, error) {
	// decrypt the token
	sigToken, err := decryptToken(encToken)
	if err != nil {
		return nil, err
	}

	// splits the signed token into the signature and the token
	separator := "\"}."
	splittedToken := strings.Split(sigToken, separator)
	// adding everything except the dot back so the JSON format is not broken
	token := splittedToken[0] + separator[:len(separator)-1]
	signature := splittedToken[1]
	if len(splittedToken) != 2 {
		return nil, &ErrInvalidToken{}
	}

	// check if the signature is valid
	if err := checkSignature(token, signature); err != nil {
		return nil, err
	}

	// create the token struct
	var tokenStruct Token
	err = json.Unmarshal([]byte(token), &tokenStruct)
	if err != nil {
		return nil, err
	}

	// check if the token is expired
	if err := checkExpiration(tokenStruct.ExpirationTimeStamp); err != nil {
		return nil, err
	}

	// check if the token is asigned to the user
	if err := checkUserAsignedToken(tokenStruct.Username, sigToken); err != nil {
		return nil, err
	}

	return &tokenStruct, nil
}

// Returns an error if provided the token is not asigned to the user (is in the database)
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

	// check if the token is the same
	if string(dbToken) != token {
		return &ErrInvalidToken{Msg: "Asigned token is not the same as the one provided"}
	}
	return nil
}

// Generates a new signature from given token and compares it to the given signature
// Returns an error if the signatures are not the same
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
