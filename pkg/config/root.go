package config

type RootConfig struct {
	ProjectID string
	Region    string

	OutputFormat string
}

type Container interface {
	GetRootConfig() *RootConfig
	GetDisplayName() string
}
