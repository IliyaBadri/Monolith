package database

import (
	"database/sql"

	"github.com/iliyabadri/monolith/cryptography"
	log "github.com/sirupsen/logrus"
)

type Admin struct {
	ID           int
	Username     string
	PasswordHash string
}

func AdminExists(username string) bool {
	currentDatabase := GetDatabase()

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM admins WHERE username = ?)`

	err := currentDatabase.QueryRow(query, username).Scan(&exists)

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"username": username,
		}).Fatal("Couldn't check if an admin exists")
	}

	return exists
}

func AddAdmin(username string, password string) {
	currentDatabase := GetDatabase()
	passwordHash, _ := cryptography.HashPassword(password)
	_, err := currentDatabase.Exec("INSERT INTO admins (username, password_hash) VALUES (?, ?)", username, passwordHash)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"username": username,
		}).Error("Couldn't add an admin to the database")
	}
}

func GetAdminPasswordHash(username string) string {
	currentDatabase := GetDatabase()

	var admin Admin

	query := `SELECT id, username, password_hash FROM admins WHERE username = ?`
	err := currentDatabase.QueryRow(query, username).Scan(&admin.ID, &admin.Username, &admin.PasswordHash)

	switch {
	case err == sql.ErrNoRows:
		return ""
	case err != nil:
		log.WithError(err).WithFields(log.Fields{
			"username": username,
		}).Fatal("Couldn't get an admin from the database")
		return ""
	default:
		return admin.PasswordHash
	}
}
