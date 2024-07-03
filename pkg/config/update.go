package config

type UpdateConfig struct {
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

	ConfigID string
}
