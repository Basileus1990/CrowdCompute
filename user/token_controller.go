// File responsible for:
//   - setting the user cookie
//   - setting the user token to the request context
//   - getting the user token from the context

package user

import (
	"context"
	"net/http"
	"time"

	"github.com/Basileus1990/CrowdCompute.git/user/user_token"
)

type userCookieName string

const UserCookieName userCookieName = "user-token"

// Sets the user cookie. The cookie expires in same time as the token
// (expiration time can differ depending on how long has the token been generated)
func SetUserCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     string(UserCookieName),
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().UTC().Add(user_token.TokenExpirationTime),
	}
	http.SetCookie(w, &cookie)
}

// Sets the user token to the request context, so it can be accessed everywhere
func SetUserContext(r *http.Request) (*http.Request, error) {
	token, err := GetTokenFromCookie(r)
	if err != nil {
		return nil, err
	}

	// Set the user token struct to the context
	ctx := context.WithValue(context.Background(), UserCookieName, token)
	r = r.WithContext(ctx)

	return r, nil
}

// Returns the user token from the cookie
//   - If the token is not found returns or is invalid an error
func GetTokenFromCookie(r *http.Request) (user_token.UserToken, error) {
	// Get the token from the cookie
	cookie, err := r.Cookie(string(UserCookieName))
	if err != nil {
		return user_token.UserToken{}, err
	}

	// Unpack the token
	token, err := user_token.UnpackToken(cookie.Value)
	if err != nil {
		return user_token.UserToken{}, err
	}

	return token, nil
}

// Returns the user token from the context
//   - If the token is not found second return value returns false
//   - If the token is found second return value returns true
func GetUserTokenCtx(r *http.Request) (user_token.UserToken, bool) {
	token, ok := r.Context().Value(UserCookieName).(user_token.UserToken)
	if !ok {
		return user_token.UserToken{}, false
	}
	return token, true
}
