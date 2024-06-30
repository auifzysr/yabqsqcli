package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"github.com/urfave/cli/v2"
	"google.golang.org/protobuf/types/known/structpb"
)

func create(cfg *config.CreateConfig) error {
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: cfg.ProjectID,
		Location:  cfg.Region,
	}
	p, err := tcs.Parent()
	if err != nil {
		return err
	}

	params, err := structpb.NewValue(cfg.Query)
	if err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	ctx := context.Background()
	m, err := client.CreateTransferConfig(
		ctx, &datatransferpb.CreateTransferConfigRequest{
			Parent: p,
			TransferConfig: &datatransferpb.TransferConfig{
				Name:         cfg.Name,
				DisplayName:  cfg.DisplayName,
				DataSourceId: "scheduled_query",
				Destination: &datatransferpb.TransferConfig_DestinationDatasetId{
					DestinationDatasetId: cfg.DestinationDataset,
				},
				Params: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"query": params,
					},
				},
				Schedule: cfg.Schedule,
				Disabled: cfg.Disabled,
			},
		},
	)
	if err != nil {
		return fmt.Errorf("creating transfer failed: parent: %s, %w", fmt.Sprintf(`projects/%s/locations/%s`,
			projectID, region,
		), err)
	}
	fmt.Printf("meta: %+v", m)
	return nil
}

func createCommand(rootCfg *config.RootConfig) *cli.Command {
	cfg := &config.CreateConfig{
		RootConfig: rootCfg,
	}

	return &cli.Command{
		Name:  "create",
		Usage: "create scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return create(cfg)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "scheduled query name",
				Destination: &cfg.Name,
			},
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
		},
	}
}
