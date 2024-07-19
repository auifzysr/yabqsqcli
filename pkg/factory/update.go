package factory

import (
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/structpb"
)

// TODO: not working
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
	}

	var fieldMaskPaths []string

	if cfg.Query != "" {
		params, err := structpb.NewValue(cfg.Query)
		if err != nil {
			return nil, fmt.Errorf("invalid params: %w", err)
		}
		tc.Params = &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"query": params,
			},
		}

		fieldMaskPaths = append(fieldMaskPaths, "params")
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
