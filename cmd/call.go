package cmd

import (
	"context"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/factory"
)

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
