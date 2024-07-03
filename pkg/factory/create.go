package factory

import (
	"fmt"
	"log"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	log.Printf("cfg: %+v", cfg)

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
		EmailPreferences: &datatransferpb.EmailPreferences{
			EnableFailureEmail: cfg.NotificationSendEmail,
		},
	}

	var schedule = &datatransferpb.ScheduleOptions{}
	if cfg.StartTime != "" {
		seconds, err := domain.TimestampSeconds(cfg.StartTime)
		if err != nil {
			return nil, err
		}
		schedule.StartTime = &timestamppb.Timestamp{
			Seconds: seconds,
		}
	}
	if cfg.EndTime != "" {
		seconds, err := domain.TimestampSeconds(cfg.EndTime)
		if err != nil {
			return nil, err
		}
		schedule.EndTime = &timestamppb.Timestamp{
			Seconds: seconds,
		}
	}
	tc.ScheduleOptions = schedule

	if cfg.NotificationPubSubTopic != "" {
		topicName, err := (&domain.PubSubTopic{
			ProjectID: cfg.ProjectID,
			TopicID:   cfg.NotificationPubSubTopic,
		}).Name()
		if err != nil {
			return nil, err

		}
		tc.NotificationPubsubTopic = topicName
	}

	req := &datatransferpb.CreateTransferConfigRequest{
		Parent:         p,
		TransferConfig: tc,
	}

	if cfg.ServiceAccountEmail != "" {
		req.ServiceAccountName = cfg.ServiceAccountEmail
	}

	return req, nil
}
