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

	"github.com/Basileus1990/CrowdCompute.git/database"
)

// TODO: change to env variable
const privateEncryptionKey = "nx/P5.,nSrqu9Owu:7vSRdSjZBP1cck!"

// For how long the token is valid in minutes
const tokenExpirationTime = 15 * time.Minute

type Token struct {
	ExpirationTimeStamp string `json:"time_stamp"`
	Username            string `json:"username"`
}

// Takes the username and generates an encrypted signed token.
//   - Generated tokens have to be added to the database.
func GenerateToken(username string) (string, error) {
	// check if the user exists
	if err := checkUser(username); err != nil {
		return "", err
	}

	// create the token
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

	// sign the token
	signedToken := fmt.Sprintf("%s.%s", tokenJSON, sig)

	// encrypt the signed token
	encryptedSigToken, err := encryptToken(signedToken)
	if err != nil {
		return "", err
	}

	return encryptedSigToken, nil
}

// returns an encrypted signed token encoded using base64
func encryptToken(sigToken string) (string, error) {
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

	encrypted := gcm.Seal(nonce, nonce, []byte(sigToken), nil)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// returns a decrypted signed token decoded from base64
func decryptToken(encodedSigToken string) (string, error) {
	decodedFromBase64, err := base64.StdEncoding.DecodeString(encodedSigToken)
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
	nonce, ciphertext := decodedFromBase64[:nonceSize], decodedFromBase64[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// creates a hash from the token
func getSignature(token []byte) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write(token)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// checks if the user exists
func checkUser(username string) error {
	exists, err := database.UserExists(username)
	if err != nil {
		return err
	}
	if !exists {
		return &ErrInvalidTokenData{Msg: "User does not exist"}
	}
	return nil
}
