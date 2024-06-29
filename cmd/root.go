package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"github.com/urfave/cli/v2"
)

var (
	projectID        string
	transferConfigID string

	region = "asia-northeast1"

	client *domain.Client
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
			&cli.StringFlag{
				Name:        "region",
				Aliases:     []string{"g"},
				Value:       "",
				Usage:       "region",
				Destination: &region,
			},
		},
		Commands: []*cli.Command{
			getCommand(),
			listCommand(),
			createCommand(),
			updateCommand(),
			deleteCommand(),
		},
	}

	var err error
	ctx := context.Background()
	client, err = domain.InitClient(ctx)
	if err != nil {
		return fmt.Errorf("data transfer client failed: %w", err)
	}

	return app.Run(os.Args)
}
