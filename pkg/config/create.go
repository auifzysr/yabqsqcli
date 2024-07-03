package config

func (*CreateConfig) configFunc() {}

type CreateConfig struct {
	*RootConfig

	DisplayName             string
	DestinationDatasetID    string
	Query                   string
	Schedule                string
	Disabled                bool
	NotificationPubSubTopic string
	NotificationSendEmail   bool
	ServiceAccountEmail     string
	StartTime               string
	EndTime                 string

	// TODO: add options available on google cloud console:

	// destinationDatasetID string
	// destinationTableID string
	// destinationTablePartitioningField string
	// destinationTableWritePreference string
	// automaticLocationSelection bool
	// locationType string
	// locationRegion string

	// TODO: implement hereafter if needed
	// encryptionKey string
}
