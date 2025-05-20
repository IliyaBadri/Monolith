package database

import (
	"database/sql"
	"sync"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

var globalDatabase *sql.DB = nil
var globalDatabaseLock sync.Mutex

var tableInitializationQueries = []string{
	`CREATE TABLE IF NOT EXISTS admins (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password_hash TEXT
	);`,
	`CREATE TABLE IF NOT EXISTS admin_sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		token TEXT UNIQUE,
		username TEXT
	);`,
}

func InitializeDatabase() {
	globalDatabaseLock.Lock()
	defer globalDatabaseLock.Unlock()
	var err error = nil
	globalDatabase, err = sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		log.WithError(err).Fatal("Database initialization failed")
	}
}

func GetDatabase() *sql.DB {
	globalDatabaseLock.Lock()
	defer globalDatabaseLock.Unlock()
	if globalDatabase == nil {
		log.Fatal("Database was accessed before being initialized")
	}
	return globalDatabase
}

func InitializeTables() {
	currentDatabase := GetDatabase()

	for _, query := range tableInitializationQueries {
		_, err := currentDatabase.Exec(query)
		if err != nil {
			log.WithError(err).Fatal("Database table initialization failed")
		}
	}
}
