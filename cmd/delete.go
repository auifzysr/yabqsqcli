package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"github.com/urfave/cli/v2"
)

func delete(cfg *config.DeleteConfig) error {
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
	err = client.DeleteTransferConfig(
		ctx, &datatransferpb.DeleteTransferConfigRequest{
			Name: n,
		},
	)
	if err != nil {
		return fmt.Errorf("getting transfer failed: %w", err)
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
