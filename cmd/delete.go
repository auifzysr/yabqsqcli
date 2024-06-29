package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/domain"
	"github.com/urfave/cli/v2"
)

func delete() error {
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

func deleteCommand() *cli.Command {
	return &cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return delete()
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
