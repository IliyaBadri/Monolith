package cryptography

import (
	"encoding/base64"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, hashingError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashingError != nil {
		return "", hashingError
	}
	base64Hash := base64.StdEncoding.EncodeToString(hash)
	return base64Hash, nil
}

func VerifyHash(base64Hash string, password string) bool {
	hash, decodingError := base64.StdEncoding.DecodeString(base64Hash)
	if decodingError != nil {
		return false
	}
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
