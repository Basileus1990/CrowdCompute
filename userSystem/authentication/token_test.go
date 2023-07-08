package authentication

import (
	"encoding/json"
	"log"
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
		"cb1I8PyfyW1ggfqipwIHcghkMsjciOxi!123$5tTT$T$ó",
		"Y4QJQz2Ln5o3wahSVijcQyYRg7OLryOh",
	}

	for _, username := range testUsernames {
		token, err := GenerateToken(username)
		if err != nil {
			t.Error(err)
		}
		decToken, err := decryptToken(token)
		if err != nil {
			t.Error(err)
		}

		// check if the username is the same after decryption
		jsonToken := strings.Split(decToken, ".")[0]
		var tokenStruct Token
		err = json.Unmarshal([]byte(jsonToken), &tokenStruct)
		if err != nil {
			log.Fatal(err)
		}
		if tokenStruct.Username != username {
			t.Errorf("Expected username %s, got %s", username, tokenStruct.Username)
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
		t.Error(err)
	}
	defer database.DeleteUser(user.Username)

	// check if the token is valid and try to asign a new one
	for i := 0; i < 2; i++ {
		token, err := GenerateToken(user.Username)
		if err != nil {
			t.Error(err)
		}
		err = database.SetAuthToken(user.Username, token)
		if err != nil {
			t.Error(err)
		}
		_, err = VerifyToken(token)
		if err != nil {
			t.Error(err)
		}
	}

	// check if the token is invalid
	notActiveToken, err := GenerateToken(user.Username)
	if err != nil {
		t.Error(err)
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
			t.Errorf("Expected error, got nil: %s", token)
		}
	}

}
