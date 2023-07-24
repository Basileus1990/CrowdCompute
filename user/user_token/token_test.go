package user_token

import (
	"testing"

	utilityTest "github.com/Basileus1990/CrowdCompute.git/utility_test"
)

const numberOfChecks = 100

func TestCreatingToken(t *testing.T) {
	for i := 0; i < numberOfChecks; i++ {
		username := utilityTest.GenerateRandomString(1000)
		_, err := GenerateToken(username)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestValidatingTokens(t *testing.T) {
	for i := 0; i < numberOfChecks; i++ {
		username := utilityTest.GenerateRandomString(1000)
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
