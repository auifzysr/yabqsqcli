package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"github.com/urfave/cli/v2"
)

func list() error {
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: projectID,
		Location:  region,
	}
	p, err := tcs.Parent()
	if err != nil {
		return err
	}
	ctx := context.Background()
	itr := client.ListTransferConfigs(ctx,
		&datatransferpb.ListTransferConfigsRequest{
			Parent: p,
		},
	)
	for {
		c, err := itr.Next()
		if err != nil {
			fmt.Printf("EOL or failed to iterate response: %s", err)
			break
		}
		fmt.Printf("config: %+v\n", c)
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
