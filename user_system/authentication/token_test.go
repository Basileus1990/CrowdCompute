package authentication

import (
	"math/rand"
	"testing"
)

func generateRandomUsername() string {
	length := rand.Intn(1000) + 1
	ran_str := make([]byte, length)

	const space int = 32
	const tilde int = 126
	for i := 0; i < length; i++ {
		ran_str[i] = byte(space + rand.Intn(tilde-space))
	}

	return string(ran_str)
}

const numberOfChecks = 100

func TestCreatingToken(t *testing.T) {
	for i := 0; i < numberOfChecks; i++ {
		username := generateRandomUsername()
		_, err := GenerateToken(username)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestValidatingTokens(t *testing.T) {
	for i := 0; i < numberOfChecks; i++ {
		username := generateRandomUsername()
		token, err := GenerateToken(username)
		if err != nil {
			t.Fatal(err)
		}

		tokenStruct, err := UnpackToken(token)
		if err != nil {
			t.Fatal(err)
		}

		// check if the token has good username
		if tokenStruct.User != username {
			t.Fatalf("Expected username %s, got %s", username, tokenStruct.User)
		}
	}
}
