package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/domain"
	"github.com/urfave/cli/v2"
)

func get() error {
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: projectID,
		Location:  region,
		ID:        transferConfigID,
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

func getCommand() *cli.Command {
	return &cli.Command{
		Name:    "get",
		Aliases: []string{"g"},
		Usage:   "get scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return get()
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "transferConfigID",
				Aliases:     []string{"c"},
				Value:       "",
				Usage:       "transferConfigID",
				Destination: &transferConfigID,
			},
		},
	}
}
