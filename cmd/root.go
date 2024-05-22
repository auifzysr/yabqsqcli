package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

var (
	projectID        string
	transferConfigID string

	region = "asia-northeast1"
)

func Run() error {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "projectID",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "projectID",
				Destination: &projectID,
			},
		},
		Commands: []*cli.Command{
			getCommand(),
			listCommand(),
			createCommand(),
			updateCommand(),
		},
	}

	return app.Run(os.Args)
}
