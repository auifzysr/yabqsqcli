// Code generated by gen/main.go; DO NOT EDIT.

package cmd

import (
	"context"
	"fmt"

	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/factory"
	"github.com/urfave/cli/v2"
)

func create(cfg *config.CreateConfig) error {
	tc, err := factory.CreateTransferConfigFactory(cfg)
	if err != nil {
		return err
	}
	ctx := context.Background()

	m, err := client.CreateTransferConfig(ctx, tc)
	if err != nil {
		return fmt.Errorf("create transfer failed: parent: %s, %w", fmt.Sprintf("projects/%s/locations/%s",
			cfg.ProjectID, cfg.Region,
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
		Name:        "create",
		Description: "create scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return create(cfg)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "name",
				Destination: &cfg.DisplayName,
			},
			&cli.StringFlag{
				Name:        "query",
				Aliases:     []string{"q"},
				Value:       "",
				Usage:       "query text",
				Destination: &cfg.Query,
			},
			&cli.StringFlag{
				Name:        "dataset",
				Aliases:     []string{"d"},
				Value:       "",
				Usage:       "destination dataset",
				Destination: &cfg.DestinationDatasetID,
			},
			&cli.StringFlag{
				Name:        "table",
				Aliases:     []string{"t"},
				Value:       "",
				Usage:       "destination table",
				Destination: &cfg.DestinationTableID,
			},
			&cli.StringFlag{
				Name:        "partitioning-field",
				Aliases:     []string{"pf"},
				Value:       "",
				Usage:       "destination table partitioning field",
				Destination: &cfg.DestinationTablePartitioningField,
			},
			&cli.StringFlag{
				Name:        "partitioning-type",
				Aliases:     []string{"pt"},
				Value:       "",
				Usage:       "destination table partitioning type",
				Destination: &cfg.DestinationTablePartitioningField,
			},
			&cli.StringFlag{
				Name:        "write-disposition",
				Aliases:     []string{"wd"},
				Value:       "",
				Usage:       "write disposition (WRITE_APPEND/ WRITE_TRUNCATE)",
				Destination: &cfg.WriteDisposition,
			},
			&cli.StringFlag{
				Name:        "schedule",
				Aliases:     []string{"s"},
				Value:       "",
				Usage:       "schedule",
				Destination: &cfg.Schedule,
			},
			&cli.BoolFlag{
				Name:        "disabled",
				Aliases:     []string{"x"},
				Usage:       "disabled",
				Destination: &cfg.Disabled,
			},
			&cli.StringFlag{
				Name:        "pubsub-topic",
				Aliases:     []string{"pst"},
				Value:       "",
				Usage:       "notification destination pubsub topic",
				Destination: &cfg.NotificationPubSubTopic,
			},
			&cli.BoolFlag{
				Name:        "enable-email",
				Aliases:     []string{"ee"},
				Usage:       "notification send email on failure",
				Destination: &cfg.NotificationSendEmail,
			},
			&cli.StringFlag{
				Name:        "service-account",
				Aliases:     []string{"sa"},
				Value:       "",
				Usage:       "runner's service account email",
				Destination: &cfg.ServiceAccountEmail,
			},
			&cli.StringFlag{
				Name:        "start-time",
				Aliases:     []string{"st"},
				Value:       "",
				Usage:       "start time",
				Destination: &cfg.StartTime,
			},
			&cli.StringFlag{
				Name:        "end-time",
				Aliases:     []string{"et"},
				Value:       "",
				Usage:       "end time",
				Destination: &cfg.EndTime,
			},
			&cli.StringFlag{
				Name:        "encryption-key-ring",
				Aliases:     []string{"kr"},
				Value:       "",
				Usage:       "encryption key ring",
				Destination: &cfg.EncryptionKeyRing,
			},
			&cli.StringFlag{
				Name:        "encryption-key",
				Aliases:     []string{"k"},
				Value:       "",
				Usage:       "encryption key",
				Destination: &cfg.EncryptionKey,
			},
		},
	}
}
