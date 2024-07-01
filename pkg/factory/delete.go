package factory

import (
	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
)

func DeleteTransferConfigFactory(cfg *config.DeleteConfig) (*datatransferpb.DeleteTransferConfigRequest, error) {
	// TODO: resolve TransferConfigID by DisplayName
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: cfg.ProjectID,
		Location:  cfg.Region,
		ID:        cfg.TransferConfigID,
	}
	n, err := tcs.Name()
	if err != nil {
		return nil, err
	}

	return &datatransferpb.DeleteTransferConfigRequest{
		Name: n,
	}, nil
}
