package main

var configs = []struct {
	Name         string
	Options      []string
	FlagTemplate string
}{
	{
		Name: "get",
		Options: []string{
			"config-id",
			"name",
		},
	},
	// {
	// 	Name: "delete",
	// 	Options: []string{
	// 		"config-id",
	// 		"name",
	// 	},
	// },
	// {
	// 	Name: "history",
	// 	Options: []string{
	// 		"config-id",
	// 		"name"},
	// },
	// {
	// 	Name:    "list",
	// 	Options: []string{},
	// },
	// {
	// 	Name: "run",
	// 	Options: []string{
	// 		"config-id",
	// 		"name",
	// 		"since",
	// 		"until",
	// 		"at",
	// 	},
	// },
	// {
	// 	Name: "update",
	// 	Options: []string{
	// 		"name",
	// 		"query",
	// 		"dataset",
	// 		"table",
	// 		"partitioning-field",
	// 		"partitioning-type",
	// 		"write-disposition",
	// 		"schedule",
	// 		"disabled",
	// 		"pubsub-topic",
	// 		"enable-email",
	// 		"service-account",
	// 		"start-time",
	// 		"end-time",
	// 		"encryption-key-ring",
	// 		"encryption-key",
	// 	},
	// },
	{
		Name: "create",
		Options: []string{
			"name",
			"query",
			"dataset",
			"table",
			"partitioning-field",
			"partitioning-type",
			"write-disposition",
			"schedule",
			"disabled",
			"pubsub-topic",
			"enable-email",
			"service-account",
			"start-time",
			"end-time",
			"encryption-key-ring",
			"encryption-key",
		},
	},
}