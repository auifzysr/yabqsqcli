package config

func (*CreateConfig) configFunc() {}

type CreateConfig struct {
	*RootConfig

	DisplayName                       string
	DestinationDatasetID              string
	Query                             string
	Schedule                          string
	Disabled                          bool
	NotificationPubSubTopic           string
	NotificationSendEmail             bool
	ServiceAccountEmail               string
	StartTime                         string
	EndTime                           string
	DestinationTableID                string
	DestinationTablePartitioningField string
	DestinationTablePartitioningType  string
	WriteDisposition                  string
	EncryptionKeyRing                 string
	EncryptionKey                     string

	// TODO: add options available on google cloud console:

	// destinationTablePartitioningField string
	// automaticLocationSelection bool
	// locationType string
	// locationRegion string

}
