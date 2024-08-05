package factory

import (
	"fmt"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.params")
	}

	if cfg.DestinationTableID != "" {
		destinationTableIDValue, err := structpb.NewValue(cfg.DestinationTableID)
		if err != nil {
			return nil, fmt.Errorf("invalid destination_table_id: %w", err)
		}
		tc.Params.Fields["destination_table_name_template"] = destinationTableIDValue
		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.params")
	}

	// TODO: not working
	// params available can be found at:
	// https://cloud.google.com/bigquery/docs/working-with-transfers#update_a_transfer
	if cfg.DestinationTablePartitioningField != "" {
		v, err := structpb.NewValue(cfg.DestinationTablePartitioningField)
		if err != nil {
			return nil, fmt.Errorf("invalid partitioning_field: %w", err)
		}
		tc.Params.Fields["partitioning_field"] = v

		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.params")
	}

	if cfg.WriteDisposition != "" {
		writeDispositionValue, err := structpb.NewValue(cfg.WriteDisposition)
		if err != nil {
			return nil, fmt.Errorf("invalid write disposition: %w", err)
		}
		tc.Params.Fields["write_disposition"] = writeDispositionValue

		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.params")
	}

	if cfg.DisplayName != "" {
		tc.DisplayName = cfg.DisplayName
		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.display_name")
	}

	if cfg.DestinationDatasetID != "" {
		tc.Destination = &datatransferpb.TransferConfig_DestinationDatasetId{
			DestinationDatasetId: cfg.DestinationDatasetID,
		}
		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.destination_dataset_id")
	}

	// TODO: updating starttime or endtime without
	// specifying the other nullifies the unspecified
	tc.ScheduleOptions = &datatransferpb.ScheduleOptions{}
	if cfg.StartTime != "" {
		seconds, err := domain.TimestampSeconds(cfg.StartTime)
		if err != nil {
			return nil, err
		}
		tc.ScheduleOptions.StartTime = &timestamppb.Timestamp{
			Seconds: seconds,
		}

		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.schedule_options")
	}
	if cfg.EndTime != "" {
		seconds, err := domain.TimestampSeconds(cfg.EndTime)
		if err != nil {
			return nil, err
		}
		tc.ScheduleOptions.EndTime = &timestamppb.Timestamp{
			Seconds: seconds,
		}

		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.schedule_options")
	}

	if cfg.Schedule != "" {
		tc.Schedule = cfg.Schedule
		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.schedule")
	}

	if cfg.NotificationPubSubTopic != "" {
		topicName, err := (&domain.PubSubTopic{
			ProjectID: cfg.ProjectID,
			TopicID:   cfg.NotificationPubSubTopic,
		}).Name()
		if err != nil {
			return nil, err

		}
		tc.NotificationPubsubTopic = topicName

		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.notification_pubsub_topic")
	}

	// TODO: yet to be tested with the Google Cloud's API
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

		fieldMaskPaths = append(fieldMaskPaths, "transfer_config.encryption_configuration")
	}

	// >> UpdateTransferConfig updates a data transfer configuration.
	// >> All fields must be set, even if they are not updated.
	req := &datatransferpb.UpdateTransferConfigRequest{}

	// TODO: not working: nothing changes
	if cfg.ServiceAccountEmail != "" {
		req.ServiceAccountName = cfg.ServiceAccountEmail
		fieldMaskPaths = append(fieldMaskPaths, "service_account_name")
	}

	// TransferConfig works as proto.Message
	fm, err := fieldmaskpb.New(req, fieldMaskPaths...)
	if err != nil {
		return nil, fmt.Errorf("invalid fieldmask: %w", err)
	}

	// >> UpdateTransferConfig updates a data transfer configuration.
	// >> All fields must be set, even if they are not updated.
	req.TransferConfig = tc
	req.UpdateMask = fm

	return req, nil
}
