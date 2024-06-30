package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"github.com/urfave/cli/v2"
)

func get(cfg *config.GetConfig) error {
	// TODO: resolve TransferConfigID by DisplayName
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: cfg.ProjectID,
		Location:  cfg.Region,
		ID:        cfg.TransferConfigID,
	}
	n, err := tcs.Name()
	if err != nil {
		return err
	}
	ctx := context.Background()
	m, err := client.GetTransferConfig(
		ctx, &datatransferpb.GetTransferConfigRequest{
			Name: n,
		},
	)
	if err != nil {
		return fmt.Errorf("getting transfer failed: %w", err)
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
