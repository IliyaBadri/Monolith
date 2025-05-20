package actions

import (
	global "github.com/iliyabadri/monolith/global"
	logrus "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func VersionAction(context *cli.Context) error {
	logrus.Infof("you are running %s v%s", global.ServerName, global.ServerVersion)
	return nil
}
