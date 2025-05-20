package validators

import "github.com/sirupsen/logrus"

func ValidateUsername(username string) bool {
	usernameLength := len(username)
	if usernameLength < 2 {
		logrus.WithField("username", username).Debug("username validation failed because the username length is less than 2")
		return false
	}

	if usernameLength > 64 {
		logrus.WithField("username", username).Debug("username validation failed because the username length is more than 64")
		return false
	}

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"

	validChars := make(map[rune]bool)
	for _, validChar := range letters + numbers {
		validChars[validChar] = true
	}

	// Disallow starting with a number
	firstChar := rune(username[0])
	for _, numberChar := range numbers {
		if firstChar == numberChar {
			logrus.WithField("username", username).Debug("username validation failed because the username starts with a number")
			return false
		}
	}

	for _, usernameChar := range username {
		if !validChars[usernameChar] {
			logrus.WithFields(logrus.Fields{
				"username":     username,
				"invalid_char": string(usernameChar),
			}).Debug("username validation failed because the username contains an invalid character")
			return false
		}
	}

	return true
}
