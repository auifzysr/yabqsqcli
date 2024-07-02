package config

func (*CreateConfig) configFunc() {}

type CreateConfig struct {
	*RootConfig

	Name                    string
	DisplayName             string
	DestinationDataset      string
	Query                   string
	Schedule                string
	Disabled                bool
	NotificationPubSubTopic string

	// TODO: add options available on google cloud console:
	// repeatFrequency string
	// repeatsEvery string
	// startNow bool
	// startAtSetTime bool
	// startDateAndRunTime time.Time
	// endNever bool
	// scheduleEndTime time.Time
	// destinationDatasetID string
	// destinationTableID string
	// destinationTablePartitioningField string
	// destinationTableWritePreference string
	// automaticLocationSelection bool
	// locationType string
	// locationRegion string
	// serviceAccountEmail string
	// encryptionKey string
	// notificationSendEmailNotifications bool
	// notificationSendCloudPubSubTopic string
}
