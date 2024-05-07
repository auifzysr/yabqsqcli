package cmd

import (
	"context"
	"fmt"

	datatransfer "cloud.google.com/go/bigquery/datatransfer/apiv1"
	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/urfave/cli/v2"
)

func list() error {
	ctx := context.Background()
	c, err := datatransfer.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("data transfer client failed: %w", err)
	}
	itr := c.ListTransferConfigs(ctx,
		&datatransferpb.ListTransferConfigsRequest{
			Parent: fmt.Sprintf(`projects/%s/locations/%s`, projectID, region),
		},
	)
	for {
		c, err := itr.Next()
		if err != nil {
			fmt.Printf("EOL or failed to iterate response: %s", err)
			break
		}
		fmt.Printf("config: %+v", c)
	}
	return nil
}

func listCommand() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list scheduled query configs",
		Action: func(cCtx *cli.Context) error {
			return list()
		},
	}
}
