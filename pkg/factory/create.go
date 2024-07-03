package factory

import (
	"fmt"
	"log"

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

	log.Printf("cfg: %+v", cfg)

	params, err := structpb.NewValue(cfg.Query)
	if err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	// TransferConfig has field "name" which is ignored when creating a transfer
	// See details on the struct's source code
	tc := &datatransferpb.TransferConfig{
		DisplayName:  cfg.DisplayName,
		DataSourceId: "scheduled_query",
		Destination: &datatransferpb.TransferConfig_DestinationDatasetId{
			DestinationDatasetId: cfg.DestinationDatasetID,
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

	req := &datatransferpb.CreateTransferConfigRequest{
		Parent:         p,
		TransferConfig: tc,
	}

	if cfg.ServiceAccountEmail != "" {
		req.ServiceAccountName = cfg.ServiceAccountEmail
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

	return req, nil
}
