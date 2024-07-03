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
		},
	}
}
