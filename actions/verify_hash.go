package actions

import (
	fmt "fmt"
	syscall "syscall"

	cryptography "github.com/iliyabadri/monolith/cryptography"
	logrus "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
	term "golang.org/x/term"
)

func VerifyHashAction(context *cli.Context) error {
	hashString := context.String("hash")
	if hashString == "" {
		return fmt.Errorf("there was no hash provided")
	}
	logrus.Info("listening for password on stdin")
	passwordBytes, readingError := term.ReadPassword(syscall.Stdin)
	if readingError != nil {
		logrus.Error("there was an error while reading the password")
		return readingError
	}
	logrus.WithFields(logrus.Fields{
		"hash": hashString,
	}).Info("checking hash")
	passwordString := string(passwordBytes)
	verified := cryptography.VerifyHash(hashString, passwordString)

	if verified {
		logrus.Info("the hash and the password match")
	} else {
		logrus.Error("the hash and the password do not match")
	}
	return nil
}
