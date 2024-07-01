package factory

import (
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"google.golang.org/protobuf/types/known/structpb"
)

func CreateTransferConfigFactory(cfg *config.CreateConfig) (*datatransferpb.CreateTransferConfigRequest, error) {
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: cfg.ProjectID,
		Location:  cfg.Region,
	}
	p, err := tcs.Parent()
	if err != nil {
		return nil, err
	}

	params, err := structpb.NewValue(cfg.Query)
	if err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	tc := &datatransferpb.TransferConfig{
		Name:         cfg.Name,
		DisplayName:  cfg.DisplayName,
		DataSourceId: "scheduled_query",
		Destination: &datatransferpb.TransferConfig_DestinationDatasetId{
			DestinationDatasetId: cfg.DestinationDataset,
		},
		Params: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"query": params,
			},
		},
		Schedule: cfg.Schedule,
		Disabled: cfg.Disabled,
	}

	req := &datatransferpb.CreateTransferConfigRequest{
		Parent:         p,
		TransferConfig: tc,
	}

	return req, nil
}