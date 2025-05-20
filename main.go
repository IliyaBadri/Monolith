package main

import (
	os "os"

	actions "github.com/iliyabadri/monolith/actions"
	logrus "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("initializing cli app")
	cliApp := &cli.App{
		Usage: "a monolithic webserver",
		Commands: []*cli.Command{
			{
				Name:   "version",
				Usage:  "shows the version of the app",
				Action: actions.VersionAction,
			},
			{
				Name:  "configure",
				Usage: "creates a configuration file",
				Flags: []cli.Flag{
					&cli.PathFlag{
						TakesFile: true,
						Name:      "output",
						Aliases:   []string{"o"},
						Usage:     "specifies the file to save the config output to",
						Required:  true,
					},
					&cli.StringFlag{
						Name:  "adminuser",
						Usage: "specifies the admin username",
					},
					&cli.StringFlag{
						Name:  "adminhash",
						Usage: "specifies the admin password hash",
					},
				},
				Action: actions.ConfigureAction,
			},
			{
				Name:   "gen-hash-interactive",
				Usage:  "generates a safe hash from a password interactively via your terminal",
				Action: actions.GenerateHashAction,
			},
			{
				Name:  "verify-hash-interactive",
				Usage: "verifies a hash with a password interactively via your terminal",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "hash",
						Aliases:  []string{"s"},
						Usage:    "specifies the password hash",
						Required: true,
					},
				},
				Action: actions.VerifyHashAction,
			},
		},
	}

	logrus.Trace("running cli app")
	err := cliApp.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
