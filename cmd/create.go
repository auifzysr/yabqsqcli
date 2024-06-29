package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"github.com/urfave/cli/v2"
	"google.golang.org/protobuf/types/known/structpb"
)

func create(config *createConfig) error {
	params, err := structpb.NewValue(config.query)
	if err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: projectID,
		Location:  region,
	}
	p, err := tcs.Parent()
	if err != nil {
		return err
	}
	ctx := context.Background()
	m, err := client.CreateTransferConfig(
		ctx, &datatransferpb.CreateTransferConfigRequest{
			Parent: p,
			TransferConfig: &datatransferpb.TransferConfig{
				Name:         config.name,
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

type createConfig struct {
	name               string
	displayName        string
	destinationDataset string
	query              string
	schedule           string
	disabled           bool

	// TODO: add options available on google cloud console:
	// repeatFrequency string
	// repeatsEvery string
	// startNow bool
	// startAtSetTime bool
	// startDateAndRunTime time.Time
	// endNever bool
	// scheduleEndTime time.Time
	// destinationDatasetID string
	// destinationTableID string
	// destinationTablePartitioningField string
	// destinationTableWritePreference string
	// automaticLocationSelection bool
	// locationType string
	// locationRegion string
	// serviceAccountEmail string
	// encryptionKey string
	// notificationSendEmailNotifications bool
	// notificationSendCloudPubSubTopic string
}

func createCommand() *cli.Command {
	config := &createConfig{}
	return &cli.Command{
		Name:  "create",
		Usage: "create scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return create(config)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "scheduled query name",
				Destination: &config.name,
			},
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
		},
	}
}
