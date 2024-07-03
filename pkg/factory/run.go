package factory

import (
	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RunTransferConfigFactory(cfg *config.RunConfig) (*datatransferpb.StartManualTransferRunsRequest, error) {
	tcs := &domain.TransferConfigsPathSpec{
		ProjectID: cfg.ProjectID,
		Location:  cfg.Region,
		ID:        cfg.TransferConfigID,
	}
	n, err := tcs.Name()
	if err != nil {
		return nil, err
	}

	req := &datatransferpb.StartManualTransferRunsRequest{
		Parent: n,
	}

	// TODO: exit with error if either is empty
	if cfg.Since != "" && cfg.Until != "" {
		s, err := domain.TimestampSeconds(cfg.Since)
		if err != nil {
			return nil, err
		}
		u, err := domain.TimestampSeconds(cfg.Until)
		if err != nil {
			return nil, err
		}
		req.Time = &datatransferpb.StartManualTransferRunsRequest_RequestedTimeRange{
			RequestedTimeRange: &datatransferpb.StartManualTransferRunsRequest_TimeRange{
				StartTime: &timestamppb.Timestamp{
					Seconds: s,
				},
				EndTime: &timestamppb.Timestamp{
					Seconds: u,
				},
			},
		}
	}

	return req, nil
}
