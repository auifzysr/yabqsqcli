package cmd

import (
	"context"
	"fmt"

	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/factory"
	"github.com/urfave/cli/v2"
)

func history(cfg *config.HistoryConfig) error {
	tc, err := factory.HistoryConfigFactory(cfg)
	if err != nil {
		return err
	}
	ctx := context.Background()

	itr := client.ListTransferRuns(ctx, tc)
	for {
		c, err := itr.Next()
		if err != nil {
			fmt.Printf("EOL or failed to iterate response: %s", err)
			break
		}
		fmt.Printf("run: %+v\n", c)
	}
	return nil
}

func historyCommand(rootCfg *config.RootConfig) *cli.Command {
	cfg := &config.HistoryConfig{
		RootConfig: rootCfg,
	}

	return &cli.Command{
		Name:    "history",
		Aliases: []string{"y"},
		Usage:   "scheduled query run history",
		Action: func(cCtx *cli.Context) error {
			return history(cfg)
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
		},
	}
}
