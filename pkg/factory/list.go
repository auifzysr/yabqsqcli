package factory

import (
	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
)

func ListTransferConfigFactory(cfg *config.ListConfig) (*datatransferpb.ListTransferConfigsRequest, error) {
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: cfg.ProjectID,
		Location:  cfg.Region,
	}
	p, err := tcs.Parent()
	if err != nil {
		return nil, err
	}

	return &datatransferpb.ListTransferConfigsRequest{
		Parent: p,
	}, nil
}
