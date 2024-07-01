package cmd

import (
	"context"
	"fmt"

	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/factory"
	"github.com/urfave/cli/v2"
)

func delete(cfg *config.DeleteConfig) error {
	tc, err := factory.DeleteTransferConfigFactory(cfg)
	if err != nil {
		return err
	}
	ctx := context.Background()

	err = client.DeleteTransferConfig(ctx, tc)
	if err != nil {
		return fmt.Errorf("deleting transfer failed: parent: %s, %w", fmt.Sprintf(`projects/%s/locations/%s`,
			cfg.ProjectID, cfg.Region,
		), err)
	}

	return nil
}

func deleteCommand(rootCfg *config.RootConfig) *cli.Command {
	cfg := &config.DeleteConfig{
		RootConfig: rootCfg,
	}

	return &cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return delete(cfg)
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
