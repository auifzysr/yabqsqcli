package main

import (
	"context"
	"fmt"
	"log"
	"os"

	datatransfer "cloud.google.com/go/bigquery/datatransfer/apiv1"
	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/urfave/cli/v2"
)

var region = "asia-northeast1"

func main() {
	var projectID, transferConfigID string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "projectID",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "projectID",
				Destination: &projectID,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "get scheduled query config",
				Action: func(cCtx *cli.Context) error {
					ctx := context.Background()
					c, err := datatransfer.NewClient(ctx)
					if err != nil {
						return fmt.Errorf("data transfer client failed: %w", err)
					}
					m, err := c.GetTransferConfig(ctx, &datatransferpb.GetTransferConfigRequest{Name: fmt.Sprintf(`projects/%s/locations/%s/transferConfigs/%s`, projectID, region, transferConfigID)})
					if err != nil {
						return fmt.Errorf("getting transfer failed: %w", err)
					}
					fmt.Printf("meta: %+v", m)
					return nil
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
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
