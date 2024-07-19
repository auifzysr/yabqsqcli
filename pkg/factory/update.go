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
		ID:        cfg.TransferConfigID,
	}
	n, err := tcs.Name()
	if err != nil {
		return nil, err
	}

	tc := &datatransferpb.TransferConfig{
		Name: n,

		// TODO: may require fill all fields of datatransferpb.TransferConfig.Params
		// even if only a part of them are supposed to be updated
		Params: &structpb.Struct{
			Fields: map[string]*structpb.Value{},
		},
	}

	var fieldMaskPaths []string

	if cfg.Query != "" {
		query, err := structpb.NewValue(cfg.Query)
		if err != nil {
			return nil, fmt.Errorf("invalid params: %w", err)
		}
		tc.Params.Fields["query"] = query

		fieldMaskPaths = append(fieldMaskPaths, "params")
	}

	if cfg.DisplayName != "" {
		tc.DisplayName = cfg.DisplayName
		fieldMaskPaths = append(fieldMaskPaths, "display_name")
	}

	if cfg.DestinationDatasetID != "" {
		tc.Destination = &datatransferpb.TransferConfig_DestinationDatasetId{
			DestinationDatasetId: cfg.DestinationDatasetID,
		}
		fieldMaskPaths = append(fieldMaskPaths, "destination_dataset_id")
	}

	if cfg.DestinationTableID != "" {
		destinationTableIDValue, err := structpb.NewValue(cfg.DestinationTableID)
		if err != nil {
			return nil, fmt.Errorf("invalid destination_table_id: %w", err)
		}
		if tc.Params == nil {
			tc.Params = &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"destination_table_name_template": destinationTableIDValue,
				},
			}
		} else {
			tc.Params.Fields["destination_table_name_template"] = destinationTableIDValue
		}
	}

	// TransferConfig works as proto.Message
	fm, err := fieldmaskpb.New(tc, fieldMaskPaths...)
	if err != nil {
		return nil, fmt.Errorf("invalid fieldmask: %w", err)
	}

	return &datatransferpb.UpdateTransferConfigRequest{
		TransferConfig: tc,
		UpdateMask:     fm,
	}, nil
}
