package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"github.com/urfave/cli/v2"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/structpb"
)

// TODO: works only with displayName
func update(cfg *config.UpdateConfig) error {
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: cfg.ProjectID,
		Location:  cfg.Region,
		ID:        cfg.ConfigID,
	}
	n, err := tcs.Name()
	if err != nil {
		return err
	}

	params, err := structpb.NewValue(cfg.Query)
	if err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	tc := &datatransferpb.TransferConfig{
		Name:         n,
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
	}
	// TransferConfig works as proto.Message
	// TODO: get specified option name and append to the tail
	fm, err := fieldmaskpb.New(tc, "params")
	if err != nil {
		return fmt.Errorf("invalid fieldmask: %w", err)
	}
	ctx := context.Background()
	m, err := client.UpdateTransferConfig(
		ctx, &datatransferpb.UpdateTransferConfigRequest{
			TransferConfig: tc,
			UpdateMask:     fm,
		},
	)
	if err != nil {
		return fmt.Errorf("updating transfer failed: name=%s, err=%w",
			fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s",
				cfg.ProjectID, cfg.Region, cfg.ConfigID), err)
	}
	fmt.Printf("meta: %+v", m)
	return nil
}

func updateCommand(rootCfg *config.RootConfig) *cli.Command {
	cfg := &config.UpdateConfig{
		RootConfig: rootCfg,
	}

	return &cli.Command{
		Name:  "update",
		Usage: "update scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return update(cfg)
		},
		Flags: []cli.Flag{
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
			&cli.StringFlag{
				Name:        "transferConfigID",
				Aliases:     []string{"c"},
				Value:       "",
				Usage:       "transferConfigID",
				Destination: &cfg.ConfigID,
			},
		},
	}
}
