package actions

import (
	syscall "syscall"

	cryptography "github.com/iliyabadri/monolith/cryptography"
	logrus "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
	term "golang.org/x/term"
)

func GenerateHashAction(context *cli.Context) error {
	logrus.Info("listening for password on stdin")
	passwordBytes, readingError := term.ReadPassword(syscall.Stdin)
	if readingError != nil {
		logrus.Error("there was an error while reading the password")
		return readingError
	}
	passwordString := string(passwordBytes)
	passwordHash, hashingError := cryptography.HashPassword(passwordString)
	if hashingError != nil {
		logrus.Error("there was an error while hashing the password")
		return hashingError
	}
	logrus.WithFields(logrus.Fields{
		"hash": passwordHash,
	}).Info("the password hash has been generated")
	return nil
}
