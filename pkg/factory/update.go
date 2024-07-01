package factory

import (
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/structpb"
)

func UpdateTransferConfigFactory(cfg *config.UpdateConfig) (*datatransferpb.UpdateTransferConfigRequest, error) {
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: cfg.ProjectID,
		Location:  cfg.Region,
		ID:        cfg.ConfigID,
	}
	n, err := tcs.Name()
	if err != nil {
		return nil, err
	}
	params, err := structpb.NewValue(cfg.Query)
	if err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	tc := &datatransferpb.TransferConfig{
		Name:         n,
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
	// TransferConfig works as proto.Message
	// TODO: get specified option name and append to the tail
	fm, err := fieldmaskpb.New(tc, "params")
	if err != nil {
		return nil, fmt.Errorf("invalid fieldmask: %w", err)
	}

	return &datatransferpb.UpdateTransferConfigRequest{
		TransferConfig: tc,
		UpdateMask:     fm,
	}, nil
}
