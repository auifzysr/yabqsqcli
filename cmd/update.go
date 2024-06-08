package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/domain"
	"github.com/urfave/cli/v2"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/structpb"
)

func update(config *updateConfig) error {
	params, err := structpb.NewValue(config.query)
	if err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: projectID,
		Location:  region,
		ID:        transferConfigID,
	}
	n, err := tcs.Name()
	if err != nil {
		return err
	}

	tc := &datatransferpb.TransferConfig{
		Name:         n,
		DisplayName:  config.displayName,
		DataSourceId: "scheduled_query",
		Destination: &datatransferpb.TransferConfig_DestinationDatasetId{
			DestinationDatasetId: config.destinationDataset,
		},
		Params: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"query": params,
			},
		},
		Schedule: config.schedule,
		Disabled: config.disabled,
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
		return fmt.Errorf("updating transfer failed: name=%s, err=%w", fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", projectID, region, config.configID), err)
	}
	fmt.Printf("meta: %+v", m)
	return nil
}

type updateConfig struct {
	displayName        string
	destinationDataset string
	query              string
	schedule           string
	disabled           bool

	configID string
}

func updateCommand() *cli.Command {
	config := &updateConfig{}
	return &cli.Command{
		Name:  "update",
		Usage: "update scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return update(config)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "displayName",
				Aliases:     []string{"dn"},
				Value:       "",
				Usage:       "scheduled query display name",
				Destination: &config.displayName,
			},
			&cli.StringFlag{
				Name:        "query",
				Aliases:     []string{"q"},
				Value:       "",
				Usage:       "scheduled query text",
				Destination: &config.query,
			},
			&cli.StringFlag{
				Name:        "destination",
				Aliases:     []string{"dd"},
				Value:       "",
				Usage:       "scheduled query destination dataset",
				Destination: &config.destinationDataset,
			},
			&cli.StringFlag{
				Name:        "schedule",
				Aliases:     []string{"sch"},
				Value:       "",
				Usage:       "scheduled query schedule",
				Destination: &config.schedule,
			},
			&cli.BoolFlag{
				Name:        "disabled",
				Aliases:     []string{"d"},
				Value:       true,
				Usage:       "scheduled query disabled",
				Destination: &config.disabled,
			},
			&cli.StringFlag{
				Name:        "transferConfigID",
				Aliases:     []string{"c"},
				Value:       "",
				Usage:       "transferConfigID",
				Destination: &config.configID,
			},
		},
	}
}
