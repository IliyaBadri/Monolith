package actions

import (
	json "encoding/json"
	fmt "fmt"
	os "os"
	filepath "path/filepath"

	global "github.com/iliyabadri/monolith/global"
	structure "github.com/iliyabadri/monolith/structure"
	validators "github.com/iliyabadri/monolith/validators"
	logrus "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func ConfigureAction(context *cli.Context) error {
	outputFilePath := context.Path("output")
	if outputFilePath == "" {
		return fmt.Errorf("there was no output file path provided")
	}
	absoluteOutputFilePath, absoluteError := filepath.Abs(outputFilePath)
	if absoluteError != nil {
		logrus.WithFields(logrus.Fields{
			"file_path": outputFilePath,
		}).Error("error while getting the absolute file path")
		return absoluteError
	}
	configuration := structure.JsonConfiguration{
		Version:           global.ServerVersion,
		Host:              "127.0.0.1",
		Port:              8080,
		MaxBodyBytes:      10 << 20, // 10 MB
		AdminUsername:     "admin",
		AdminPasswordHash: "JDJhJDEwJER1OFJOUnFqclhVQlg0N0p1eGpBNHUxM0NHRzJEak5JN1Q3endHZ2tFR3pjc1k0eFg2ZzVT", //password (literally)
	}
	inputAdminUsername := context.String("adminuser")
	if inputAdminUsername != "" {
		isAdminUsernameValid := validators.ValidateUsername(inputAdminUsername)
		if !isAdminUsernameValid {
			return fmt.Errorf("the input admin username is not a valid username")
		}
		configuration.AdminUsername = inputAdminUsername
	}
	inputAdminPasswordHash := context.String("adminhash")
	if inputAdminPasswordHash != "" {
		configuration.AdminPasswordHash = inputAdminPasswordHash
	}
	configurationData, marshallingError := json.Marshal(configuration)
	if marshallingError != nil {
		logrus.WithFields(logrus.Fields{
			"config": configuration,
		}).Error("error while marshalling the config into json")
		return marshallingError
	}
	fileError := createAndWriteFile(absoluteOutputFilePath, configurationData)
	if fileError != nil {
		logrus.WithFields(logrus.Fields{
			"file_path": absoluteOutputFilePath,
		}).Error("error while creating the config file")
		return fileError
	}
	logrus.WithFields(logrus.Fields{
		"file_path": absoluteOutputFilePath,
	}).Info("successfully created the config file")
	return nil
}

func createAndWriteFile(absoluteFilePath string, data []byte) error {
	file, openFileErr := os.OpenFile(absoluteFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if openFileErr != nil {
		logrus.WithFields(logrus.Fields{
			"file_path": absoluteFilePath,
		}).Error("error while opening a file")
		return openFileErr
	}

	_, writeErr := file.Write(data)
	if writeErr != nil {
		logrus.WithFields(logrus.Fields{
			"file_path": absoluteFilePath,
		}).Error("error while writing to a file")

		closeErr := file.Close()
		if closeErr != nil {
			logrus.WithFields(logrus.Fields{
				"file_path": absoluteFilePath,
			}).Error("error while closing a file")
			return closeErr
		}

		return writeErr
	}

	closeErr := file.Close()
	if closeErr != nil {
		logrus.WithFields(logrus.Fields{
			"file_path": absoluteFilePath,
		}).Error("error while closing a file")
		return closeErr
	}

	return nil
}
