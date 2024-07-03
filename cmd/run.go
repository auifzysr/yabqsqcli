package cmd

import (
	"context"
	"fmt"

	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/factory"
	"github.com/urfave/cli/v2"
)

func run(cfg *config.RunConfig) error {
	tc, err := factory.RunTransferConfigFactory(cfg)
	if err != nil {
		return err
	}
	ctx := context.Background()

	m, err := client.StartManualTransferRuns(ctx, tc)
	if err != nil {
		return fmt.Errorf("running transfer failed: parent: %s, %w", fmt.Sprintf(`projects/%s/locations/%s`,
			cfg.ProjectID, cfg.Region,
		), err)
	}
	fmt.Printf("meta: %+v", m)

	return nil
}

func runCommand(rootCfg *config.RootConfig) *cli.Command {
	cfg := &config.RunConfig{
		RootConfig: rootCfg,
	}

	return &cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "run scheduled query",
		Action: func(cCtx *cli.Context) error {
			return run(cfg)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config-id",
				Aliases:     []string{"c"},
				Value:       "",
				Usage:       "transfer config ID",
				Destination: &cfg.TransferConfigID,
			},
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "name",
				Destination: &cfg.DisplayName,
			},
			&cli.StringFlag{
				Name:        "since",
				Aliases:     []string{"s"},
				Value:       "",
				Usage:       "since what past time running scheduled query, must be with --until in RFC3339",
				Destination: &cfg.Since,
			},
			&cli.StringFlag{
				Name:        "until",
				Aliases:     []string{"u"},
				Value:       "",
				Usage:       "until what past time running scheduled query, must be with --since in RFC3339",
				Destination: &cfg.Until,
			},
			&cli.StringFlag{
				Name:        "at",
				Aliases:     []string{"t"},
				Value:       "",
				Usage:       "at what time running scheduled query, in RFC3339",
				Destination: &cfg.At,
			},
		},
	}
}
