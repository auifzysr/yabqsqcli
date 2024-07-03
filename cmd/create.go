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
		return fmt.Errorf("creating transfer failed: parent: %s, %w", fmt.Sprintf(`projects/%s/locations/%s`,
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
		Name:  "create",
		Usage: "create scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return create(cfg)
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
				Name:        "destinationDatasetID",
				Aliases:     []string{"dd"},
				Value:       "",
				Usage:       "scheduled query destination dataset ID",
				Destination: &cfg.DestinationDatasetID,
			},
			&cli.StringFlag{
				Name:        "destinationTableID",
				Aliases:     []string{"dt"},
				Value:       "",
				Usage:       "scheduled query destination table ID",
				Destination: &cfg.DestinationTableID,
			},
			&cli.StringFlag{
				Name:        "writeDisposition",
				Aliases:     []string{"wd"},
				Value:       "",
				Usage:       "write disposition",
				Destination: &cfg.WriteDisposition,
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
				Value:       false,
				Usage:       "scheduled query disabled",
				Destination: &cfg.Disabled,
			},
			&cli.StringFlag{
				Name:        "pubsubtopic",
				Aliases:     []string{"pt"},
				Value:       "",
				Usage:       "notification destination pubsub topic",
				Destination: &cfg.NotificationPubSubTopic,
			},
			&cli.BoolFlag{
				Name:        "failureEmail",
				Aliases:     []string{"em"},
				Value:       false,
				Usage:       "notification send email on failure (true/false)",
				Destination: &cfg.NotificationSendEmail,
			},
			&cli.StringFlag{
				Name:        "serviceaccount",
				Aliases:     []string{"sa"},
				Value:       "",
				Usage:       "service account email to run scheduled query",
				Destination: &cfg.ServiceAccountEmail,
			},
			&cli.StringFlag{
				Name:        "startTime",
				Aliases:     []string{"st"},
				Value:       "",
				Usage:       "start time",
				Destination: &cfg.StartTime,
			},
			&cli.StringFlag{
				Name:        "endTime",
				Aliases:     []string{"et"},
				Value:       "",
				Usage:       "end time",
				Destination: &cfg.EndTime,
			},
			&cli.StringFlag{
				Name:        "encryptionKeyRing",
				Aliases:     []string{"kr"},
				Value:       "",
				Usage:       "encryption key ring",
				Destination: &cfg.EncryptionKeyRing,
			},
			&cli.StringFlag{
				Name:        "encryptionKey",
				Aliases:     []string{"k"},
				Value:       "",
				Usage:       "encryption key",
				Destination: &cfg.EncryptionKey,
			},
		},
	}
}
