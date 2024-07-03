package config

type RunConfig struct {
	*RootConfig

	TransferConfigID string
	DisplayName      string
	Since            string
	Until            string
}
