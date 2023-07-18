package usersystem

import (
	"context"
	"errors"
	"net/http"

	"github.com/Basileus1990/CrowdCompute.git/user_system/authentication"
)

type userCookieName string

const UserCookieName userCookieName = "user"

func SetUserContext(r *http.Request) (*http.Request, error) {
	cookie, err := r.Cookie(string(UserCookieName))
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			return r, nil
		default:
			return nil, err
		}
	}

	token, err := authentication.VerifyToken(cookie.Value)
	if err != nil {
		return nil, err
	}

	ctx := context.WithValue(context.Background(), UserCookieName, token.Username)
	r = r.WithContext(ctx)

	return r, nil
}

func GetUsernameCtx(r *http.Request) string {
	username, ok := r.Context().Value(UserCookieName).(string)
	if !ok {
		return ""
	}
	return username
}
