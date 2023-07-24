// File used to generate tokens

package user_token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// TODO: change to env variable
const privateSecredKey = "nx/P5.,nSrqu9Owu:7vSRdSjZBP1cck!"

// For how long the token is valid in minutes
// TODO: change to config file
const TokenExpirationTime = 15 * time.Minute

type UserToken struct {
	User string `json:"user"`
}

// Takes the username and generates a JWT token.
// Claims:
//   - user: username
//   - exp:  expiration time
func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().UTC().Add(TokenExpirationTime).Unix(),
		"user": username,
	})

	tokenString, err := token.SignedString([]byte(privateSecredKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Takes the token string and generates a new token with the same username but extended expiration time
func RenewToken(tokenString string) (string, error) {
	token, err := UnpackToken(tokenString)
	if err != nil {
		return "", err
	}

	return GenerateToken(token.User)
}

// Unpacks the token and returns the UserToken struct
//   - checks if the token is valid
//   - expiration is checked by JWT package
//
// If is not valid or has expired returns an ErrInvalidToken error
func UnpackToken(tokenString string) (UserToken, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateSecredKey), nil
	})
	if err != nil {
		return UserToken{}, err
	}
	if !token.Valid {
		return UserToken{}, &ErrInvalidToken{}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return UserToken{}, &ErrInvalidToken{}
	}

	tokenStruct := UserToken{
		User: claims["user"].(string),
	}

	return tokenStruct, nil
}
