package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"github.com/urfave/cli/v2"
)

var (
	projectID        string
	transferConfigID string

	defaultRegion = "asia-northeast1"
	region        = defaultRegion

	client *domain.Client
)

func Run() error {
	rootCfg := &config.RootConfig{
		Region: defaultRegion,
	}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "projectID",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "projectID",
				Destination: &rootCfg.ProjectID,
			},
			&cli.StringFlag{
				Name:        "region",
				Aliases:     []string{"g"},
				Value:       "",
				Usage:       "region",
				Destination: &rootCfg.Region,
			},
		},
		Commands: []*cli.Command{
			getCommand(rootCfg),
			listCommand(rootCfg),
			createCommand(rootCfg),
			updateCommand(),
			deleteCommand(rootCfg),
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
