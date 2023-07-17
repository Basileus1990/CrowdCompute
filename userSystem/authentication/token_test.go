package authentication

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
	"github.com/Basileus1990/CrowdCompute.git/database"
)

func TestCreatingToken(t *testing.T) {
	testUsernames := []string{
		"test",
		"test1",
		"lizsdufhioioyhfpoih",
		"nxuP5BVnSrqu9OwuD7vSRdSjZBP1cckB",
		"cbdfg1I8PyfyW1ggpwIHcghkMsjciOxi!123ijoij",
		"Y4QJQz2Ln5o3wahSVijcQyYRg7OLryOh",
	}

	for _, username := range testUsernames {
		// adding user for test
		user := dataStructures.User{
			Username: username,
			Email:    "test@test.test" + username,
		}
		pass := "test"
		err := database.AddUser(user, pass)
		if err != nil {
			t.Fatal(err)
		}
		defer database.DeleteUser(user.Username)

		token, err := GenerateToken(username)
		if err != nil {
			t.Fatal(err)
		}
		decToken, err := decryptToken(token)
		if err != nil {
			t.Fatal(err)
		}

		// check if the username is the same after decryption
		jsonToken := strings.Split(decToken, ".")[0]
		var tokenStruct Token
		err = json.Unmarshal([]byte(jsonToken), &tokenStruct)
		if err != nil {
			t.Fatal(err)
		}
		if tokenStruct.Username != username {
			t.Fatalf("Expected username %s, got %s", username, tokenStruct.Username)
		}
	}
}

func TestValidatingTokens(t *testing.T) {
	user := dataStructures.User{
		Username: "test",
		Email:    "test@test.test",
	}
	pass := "test"
	err := database.AddUser(user, pass)
	if err != nil {
		t.Fatal(err)
	}
	defer database.DeleteUser(user.Username)

	// check if the token is valid and try to asign a new one
	for i := 0; i < 2; i++ {
		token, err := GenerateToken(user.Username)
		if err != nil {
			t.Fatal(err)
		}
		err = database.SetAuthToken(user.Username, token)
		if err != nil {
			t.Fatal(err)
		}
		tokenStruct, err := VerifyToken(token)
		if err != nil {
			t.Fatal(err)
		}

		// check if the token has good username
		if tokenStruct.Username != user.Username {
			t.Fatalf("Expected username %s, got %s", user.Username, tokenStruct.Username)
		}
	}

	// check if the token is invalid
	notActiveToken, err := GenerateToken(user.Username)
	if err != nil {
		t.Fatal(err)
	}
	invalidTokens := []string{
		"invalidddddddddddddddddddddddddddddddddddddddddddd",
		"invalid.invaliddddddddddddddddddddddddddddddddddddd",
		"{dsfgfafg:sadgfasf}.dsaoidgiosjuopa",
		notActiveToken,
	}
	for _, token := range invalidTokens {
		_, err := VerifyToken(token)
		if err == nil {
			t.Fatalf("Expected error, got nil: %s", token)
		}
	}

}
