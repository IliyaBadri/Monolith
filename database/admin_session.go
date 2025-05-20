package database

import (
	"github.com/iliyabadri/monolith/cryptography"
	log "github.com/sirupsen/logrus"
)

func AddAdminSession(username string) string {
	currentDatabase := GetDatabase()
	sessionToken := cryptography.GenerateToken(32)
	_, err := currentDatabase.Exec("INSERT INTO admin_sessions (token, username) VALUES (?, ?)", sessionToken, username)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"username": username,
		}).Error("Couldn't add an admin session to the database")
		return ""
	}
	return sessionToken
}
