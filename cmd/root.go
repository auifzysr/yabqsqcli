package cmd

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"github.com/auifzysr/yabqsqcli/pkg/factory"
	"github.com/urfave/cli/v2"
)

var (
	defaultRegion = "asia-northeast1"
	region        = defaultRegion

	client *domain.Client
)

func Run() error {
	rootCfg := &config.RootConfig{
		Region: defaultRegion,
	}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "project",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "project ID",
				Destination: &rootCfg.ProjectID,
			},
			&cli.StringFlag{
				Name:        "region",
				Aliases:     []string{"g"},
				Value:       "",
				Usage:       "region",
				Destination: &rootCfg.Region,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Value:       "json",
				Usage:       "output format (json/plain)",
				Destination: &rootCfg.OutputFormat,
			},
		},
		Commands: []*cli.Command{
			getCommand(rootCfg),
			listCommand(rootCfg),
			createCommand(rootCfg),
			updateCommand(rootCfg),
			deleteCommand(rootCfg),
			runCommand(rootCfg),
			historyCommand(rootCfg),
		},
	}

	var err error
	ctx := context.Background()
	client, err = domain.InitClient(ctx)
	if err != nil {
		return fmt.Errorf("data transfer client failed: %w", err)
	}

	return app.Run(os.Args)
}

func callList(ctx context.Context, cfg *config.ListConfig) ([]*datatransferpb.TransferConfig, error) {
	tc, err := factory.ListTransferConfigFactory(cfg)
	if err != nil {
		return nil, err
	}
	var res []*datatransferpb.TransferConfig
	itr := client.ListTransferConfigs(ctx, tc)
	for {
		m, err := itr.Next()
		if err != nil {
			break
		}
		res = append(res, m)
	}

	return res, nil
}
