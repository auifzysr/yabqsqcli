package cmd

import (
	"context"
	"fmt"

	datatransfer "cloud.google.com/go/bigquery/datatransfer/apiv1"
	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/urfave/cli/v2"
)

func delete() error {
	ctx := context.Background()
	c, err := datatransfer.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("data transfer client failed: %w", err)
	}
	m := c.DeleteTransferConfig(
		ctx, &datatransferpb.DeleteTransferConfigRequest{
			Name: fmt.Sprintf(`projects/%s/locations/%s/transferConfigs/%s`,
				projectID, region, transferConfigID,
			)})
	if err != nil {
		return fmt.Errorf("deleteting transfer failed: %w", err)
	}
	fmt.Printf("meta: %+v", m)
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
