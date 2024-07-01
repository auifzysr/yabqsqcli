package cmd

import (
	"context"
	"fmt"

	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/factory"
	"github.com/urfave/cli/v2"
)

func get(cfg *config.GetConfig) error {
	tc, err := factory.GetTransferConfigFactory(cfg)
	if err != nil {
		return err
	}
	ctx := context.Background()

	m, err := client.GetTransferConfig(ctx, tc)
	if err != nil {
		return fmt.Errorf("getting transfer failed: parent: %s, %w", fmt.Sprintf(`projects/%s/locations/%s`,
			cfg.ProjectID, cfg.Region,
		), err)
	}
	fmt.Printf("meta: %+v", m)

	return nil
}

func getCommand(rootCfg *config.RootConfig) *cli.Command {
	cfg := &config.GetConfig{
		RootConfig: rootCfg,
	}

	return &cli.Command{
		Name:    "get",
		Aliases: []string{"g"},
		Usage:   "get scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return get(cfg)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "transferConfigID",
				Aliases:     []string{"c"},
				Value:       "",
				Usage:       "transferConfigID",
				Destination: &cfg.TransferConfigID,
			},
			&cli.StringFlag{
				Name:        "displayName",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "displayName",
				Destination: &cfg.DisplayName,
			},
		},
	}
}
