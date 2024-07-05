package main

type param struct {
	name                    string
	fieldDefinitionTemplate string
	flagDefinitionTemplate  string
}

var params = map[string]param{
	"name": {
		name:                    "name",
		fieldDefinitionTemplate: "DisplayName string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "name",
				Destination: &cfg.DisplayName,
			},`,
	},
	"config-id": {
		name:                    "config-id",
		fieldDefinitionTemplate: "TransferConfigID string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "config-id",
				Aliases:     []string{"c"},
				Value:       "",
				Usage:       "transfer config ID",
				Destination: &cfg.TransferConfigID,
			},`,
	},
	"since": {
		name:                    "since",
		fieldDefinitionTemplate: "Since string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "since",
				Aliases:     []string{"s"},
				Value:       "",
				Usage:       "since what past time running scheduled query, must be with --until in RFC3339",
				Destination: &cfg.Since,
			},`,
	},
	"until": {
		name:                    "until",
		fieldDefinitionTemplate: "Until string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "until",
				Aliases:     []string{"u"},
				Value:       "",
				Usage:       "until what past time running scheduled query, must be with --since in RFC3339",
				Destination: &cfg.Until,
			},`,
	},
	"at": {
		name:                    "at",
		fieldDefinitionTemplate: "At string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "at",
				Aliases:     []string{"a"},
				Value:       "",
				Usage:       "at what time running scheduled query, in RFC3339",
				Destination: &cfg.At,
			},`,
	},
	"query": {
		name:                    "query",
		fieldDefinitionTemplate: "Query string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "query",
				Aliases:     []string{"q"},
				Value:       "",
				Usage:       "query text",
				Destination: &cfg.Query,
			},`,
	},
	"dataset": {
		name:                    "dataset",
		fieldDefinitionTemplate: "DestinationDatasetID string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "dataset",
				Aliases:     []string{"d"},
				Value:       "",
				Usage:       "destination dataset",
				Destination: &cfg.DestinationDatasetID,
			},`,
	},
	"table": {
		name:                    "table",
		fieldDefinitionTemplate: "DestinationTableID string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "table",
				Aliases:     []string{"t"},
				Value:       "",
				Usage:       "destination table",
				Destination: &cfg.DestinationTableID,
			},`,
	},
	"partitioning-field": {
		name:                    "partitioning-field",
		fieldDefinitionTemplate: "DestinationTablePartitioningField string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "partitioning-field",
				Aliases:     []string{"pf"},
				Value:       "",
				Usage:       "destination table partitioning field",
				Destination: &cfg.DestinationTablePartitioningField,
			},`,
	},
	"partitioning-type": {
		name:                    "partitioning-type",
		fieldDefinitionTemplate: "DestinationTablePartitioningType string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "partitioning-type",
				Aliases:     []string{"pt"},
				Value:       "",
				Usage:       "destination table partitioning type",
				Destination: &cfg.DestinationTablePartitioningField,
			},`,
	},
	"write-disposition": {
		name:                    "write-disposition",
		fieldDefinitionTemplate: "WriteDisposition string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "write-disposition",
				Aliases:     []string{"wd"},
				Value:       "",
				Usage:       "write disposition (WRITE_APPEND/ WRITE_TRUNCATE)",
				Destination: &cfg.WriteDisposition,
			},`,
	},
	"schedule": {
		name:                    "schedule",
		fieldDefinitionTemplate: "Schedule string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "schedule",
				Aliases:     []string{"s"},
				Value:       "",
				Usage:       "schedule",
				Destination: &cfg.Schedule,
			},`,
	},
	"disabled": {
		name:                    "disabled",
		fieldDefinitionTemplate: "Disabled bool",
		flagDefinitionTemplate: `&cli.BoolFlag{
				Name:        "disabled",
				Aliases:     []string{"x"},
				Usage:       "disabled",
				Destination: &cfg.Disabled,
			},`,
	},
	"pubsub-topic": {
		name:                    "pubsub-topic",
		fieldDefinitionTemplate: "NotificationPubSubTopic string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "pubsub-topic",
				Aliases:     []string{"pst"},
				Value:       "",
				Usage:       "notification destination pubsub topic",
				Destination: &cfg.NotificationPubSubTopic,
			},`,
	},
	"enable-email": {
		name:                    "enable-email",
		fieldDefinitionTemplate: "NotificationSendEmail bool",
		flagDefinitionTemplate: `&cli.BoolFlag{
				Name:        "enable-email",
				Aliases:     []string{"ee"},
				Usage:       "notification send email on failure",
				Destination: &cfg.NotificationSendEmail,
			},`,
	},
	"service-account": {
		name:                    "service-account",
		fieldDefinitionTemplate: "ServiceAccountEmail string",
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "service-account",
				Aliases:     []string{"sa"},
				Value:       "",
				Usage:       "runner's service account email",
				Destination: &cfg.ServiceAccountEmail,
			},`,
	},
	"start-time": {
		name:                    "start-time",
		fieldDefinitionTemplate: `StartTime string`,
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "start-time",
				Aliases:     []string{"st"},
				Value:       "",
				Usage:       "start time",
				Destination: &cfg.StartTime,
			},`,
	},
	"end-time": {
		name:                    "end-time",
		fieldDefinitionTemplate: `EndTime string`,
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "end-time",
				Aliases:     []string{"et"},
				Value:       "",
				Usage:       "end time",
				Destination: &cfg.EndTime,
			},`,
	},
	"encryption-key-ring": {
		name:                    "encryption-key-ring",
		fieldDefinitionTemplate: `EncryptionKeyRing string`,
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "encryption-key-ring",
				Aliases:     []string{"kr"},
				Value:       "",
				Usage:       "encryption key ring",
				Destination: &cfg.EncryptionKeyRing,
			},`,
	},
	"encryption-key": {
		name:                    "encryption-key",
		fieldDefinitionTemplate: `EncryptionKey string`,
		flagDefinitionTemplate: `&cli.StringFlag{
				Name:        "encryption-key",
				Aliases:     []string{"k"},
				Value:       "",
				Usage:       "encryption key",
				Destination: &cfg.EncryptionKey,
			},`,
	},
}
