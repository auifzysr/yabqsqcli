package cmd

import (
	"context"
	"fmt"

	datatransfer "cloud.google.com/go/bigquery/datatransfer/apiv1"
	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/urfave/cli/v2"
)

func get() error {
	ctx := context.Background()
	c, err := datatransfer.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("data transfer client failed: %w", err)
	}
	m, err := c.GetTransferConfig(
		ctx, &datatransferpb.GetTransferConfigRequest{
			Name: fmt.Sprintf(`projects/%s/locations/%s/transferConfigs/%s`,
				projectID, region, transferConfigID,
			)})
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
