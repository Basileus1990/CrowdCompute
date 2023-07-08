// File used to generate tokens

package authentication

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// TODO: change to env variable
const privateEncryptionKey = "nx/P5.,nSrqu9Owu:7vSRdSjZBP1cck!"

// For how long the token is valid in minutes
const tokenExpirationTime = 15 * time.Minute

type Token struct {
	ExpirationTimeStamp string `json:"time_stamp"`
	Username            string `json:"username"`
}

// takes the username and generates a token
// username has to be valid
func GenerateToken(username string) (string, error) {
	token := Token{
		fmt.Sprint(time.Now().UnixNano() + int64(tokenExpirationTime.Nanoseconds())),
		username,
	}
	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	// create the signature
	sig, err := getSignature(tokenJSON)
	if err != nil {
		return "", err
	}

	signedToken := fmt.Sprintf("%s.%s", tokenJSON, sig)

	// encrypt the token
	encryptedToken, err := encryptToken(signedToken)
	if err != nil {
		return "", err
	}

	return encryptedToken, nil
}

func encryptToken(token string) (string, error) {
	c, err := aes.NewCipher([]byte(privateEncryptionKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encrypted := gcm.Seal(nonce, nonce, []byte(token), nil)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func decryptToken(encodedToken string) (string, error) {
	decodedFromBase64, err := base64.StdEncoding.DecodeString(encodedToken)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher([]byte(privateEncryptionKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := []byte(decodedFromBase64)[:nonceSize], []byte(decodedFromBase64)[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func getSignature(token []byte) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write(token)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
