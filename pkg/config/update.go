package config

type UpdateConfig struct {
	*RootConfig

	DisplayName        string
	DestinationDataset string
	Query              string
	Schedule           string
	Disabled           bool

	ConfigID string
}
