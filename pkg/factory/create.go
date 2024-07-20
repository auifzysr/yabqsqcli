package factory

import (
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

	// TransferConfig has field "name" which is ignored when creating a transfer
	// See details on the struct's source code
	tc := &datatransferpb.TransferConfig{
		DataSourceId: "scheduled_query",
		Params: &structpb.Struct{
			Fields: map[string]*structpb.Value{},
		},
	}

	// required
	if cfg.Query == "" {
		return nil, fmt.Errorf("query required")
	}
	query, err := structpb.NewValue(cfg.Query)
	if err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	tc.Params.Fields["query"] = query

	// optional hereafter
	if tc.DisplayName != "" {
		tc.DisplayName = cfg.DisplayName
	}
	if cfg.DestinationDatasetID != "" {
		tc.Destination = &datatransferpb.TransferConfig_DestinationDatasetId{
			DestinationDatasetId: cfg.DestinationDatasetID,
		}
	}
	if cfg.Schedule != "" {
		tc.Schedule = cfg.Schedule
	}
	if cfg.Disabled {
		tc.Disabled = cfg.Disabled
	}
	if cfg.NotificationSendEmail != "" {
		tc.EmailPreferences = &datatransferpb.EmailPreferences{
			EnableFailureEmail: cfg.NotificationSendEmail,
		}
	}

	// params available can be found at:
	// https://cloud.google.com/bigquery/docs/working-with-transfers#update_a_transfer
	if cfg.DestinationTablePartitioningField != "" && cfg.DestinationTablePartitioningType != "" {
		{
			v, err := structpb.NewValue(cfg.DestinationTablePartitioningField)
			if err != nil {
				return nil, fmt.Errorf("invalid partitioning_field: %w", err)
			}
			tc.Params.Fields["partitioning_field"] = v
		}

		{
			v, err := structpb.NewValue(cfg.DestinationTablePartitioningType)
			if err != nil {
				return nil, fmt.Errorf("invalid partitioning_type: %w", err)
			}
			tc.Params.Fields["partitioning_type"] = v
		}
	}

	if cfg.DestinationTableID != "" {
		destinationTableIDValue, err := structpb.NewValue(cfg.DestinationTableID)
		if err != nil {
			return nil, fmt.Errorf("invalid destination_table_id: %w", err)
		}
		tc.Params.Fields["destination_table_name_template"] = destinationTableIDValue
	}

	if cfg.WriteDisposition != "" {
		writeDispositionValue, err := structpb.NewValue(cfg.WriteDisposition)
		if err != nil {
			return nil, fmt.Errorf("invalid write disposition: %w", err)
		}
		tc.Params.Fields["write_disposition"] = writeDispositionValue
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

	if cfg.EncryptionKeyRing != "" && cfg.EncryptionKey != "" {
		k, err := (&domain.KMS{
			ProjectID: cfg.ProjectID,
			Location:  cfg.Region,
			KeyRing:   cfg.EncryptionKeyRing,
			Key:       cfg.EncryptionKey,
		}).ResourceID()
		if err != nil {
			return nil, err
		}
		ec := &datatransferpb.EncryptionConfiguration{
			KmsKeyName: &wrapperspb.StringValue{
				Value: k,
			},
		}
		tc.EncryptionConfiguration = ec
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
