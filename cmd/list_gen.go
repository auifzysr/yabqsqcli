// Code generated by gen/main.go; DO NOT EDIT.

package cmd

import (
	"context"
	"fmt"

	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/factory"
	"github.com/urfave/cli/v2"
)

func list(cfg *config.ListConfig) error {
	tc, err := factory.ListTransferConfigFactory(cfg)
	if err != nil {
		return err
	}
	ctx := context.Background()

	itr := client.ListTransferConfigs(ctx, tc)
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

func listCommand(rootCfg *config.RootConfig) *cli.Command {
	cfg := &config.ListConfig{
		RootConfig: rootCfg,
	}

	return &cli.Command{
		Name:        "list",
		Description: "list scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return list(cfg)
		},
		Flags: []cli.Flag{},
	}
}
