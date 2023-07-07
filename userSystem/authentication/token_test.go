package authentication

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
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
