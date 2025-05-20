package cryptography

import (
	"crypto/rand"
	"encoding/hex"

	log "github.com/sirupsen/logrus"
)

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"length": length,
		}).Fatal("Couldn't generate a token")
		return ""
	}
	return hex.EncodeToString(bytes)
}
