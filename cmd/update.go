package cmd

import (
	"context"
	"fmt"

	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/factory"
	"github.com/urfave/cli/v2"
)

// TODO: works only with displayName
func update(cfg *config.UpdateConfig) error {
	tc, err := factory.UpdateTransferConfigFactory(cfg)
	if err != nil {
		return err
	}
	ctx := context.Background()

	m, err := client.UpdateTransferConfig(ctx, tc)
	if err != nil {
		return fmt.Errorf("updating transfer failed: name=%s, err=%w",
			fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s",
				cfg.ProjectID, cfg.Region, cfg.ConfigID), err)
	}
	fmt.Printf("meta: %+v", m)

	return nil
}

func updateCommand(rootCfg *config.RootConfig) *cli.Command {
	cfg := &config.UpdateConfig{
		RootConfig: rootCfg,
	}

	return &cli.Command{
		Name:  "update",
		Usage: "update scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return update(cfg)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "displayName",
				Aliases:     []string{"dn"},
				Value:       "",
				Usage:       "scheduled query display name",
				Destination: &cfg.DisplayName,
			},
			&cli.StringFlag{
				Name:        "query",
				Aliases:     []string{"q"},
				Value:       "",
				Usage:       "scheduled query text",
				Destination: &cfg.Query,
			},
			&cli.StringFlag{
				Name:        "destination",
				Aliases:     []string{"dd"},
				Value:       "",
				Usage:       "scheduled query destination dataset",
				Destination: &cfg.DestinationDataset,
			},
			&cli.StringFlag{
				Name:        "schedule",
				Aliases:     []string{"sch"},
				Value:       "",
				Usage:       "scheduled query schedule",
				Destination: &cfg.Schedule,
			},
			&cli.BoolFlag{
				Name:        "disabled",
				Aliases:     []string{"d"},
				Value:       true,
				Usage:       "scheduled query disabled",
				Destination: &cfg.Disabled,
			},
			&cli.StringFlag{
				Name:        "transferConfigID",
				Aliases:     []string{"c"},
				Value:       "",
				Usage:       "transferConfigID",
				Destination: &cfg.ConfigID,
			},
		},
	}
}
