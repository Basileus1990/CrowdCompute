package user

import (
	"net/http"
	"testing"
	"time"

	"github.com/Basileus1990/CrowdCompute.git/user/user_token"
	utilitytest "github.com/Basileus1990/CrowdCompute.git/utility_test"
)

const numberOfChecks = 100

// Creates a new user token, asigns it to a request's cookie,
// then inserts it into the context
// and checks if the token is found in the context
func TestTokenCTX(t *testing.T) {
	for i := 0; i < numberOfChecks; i++ {
		// Create a new token
		username := utilitytest.GenerateRandomString(1000)
		token, err := user_token.GenerateToken(username)
		if err != nil {
			t.Fatal(err)
		}

		// Create a new request with the token in a cookie
		// Not using SetUserCookie because it needs a responseWriter
		cookie := http.Cookie{
			Name:     string(UserCookieName),
			Value:    token,
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Now().UTC().Add(user_token.TokenExpirationTime),
		}
		header := http.Header{}
		header.Add("Cookie", cookie.String())
		request := &http.Request{Header: header}

		// Set the token to the context
		r, err := SetUserContext(request)
		if err != nil {
			t.Fatal(err)
		}

		// Check if the token is found in the context
		userToken, ok := GetUserTokenCtx(r)
		if !ok {
			t.Fatal("Token not found in context")
		}
		if userToken.User != username {
			t.Fatalf("Expected username %s, got %s", username, userToken.User)
		}
	}
}
