package main

var scenarios = []struct {
	Name                    string
	CallResolver            bool
	IsConfig                bool
	IsItr                   bool
	IsSlice                 bool
	Options                 []string
	APICallFuncName         string
	APICallResponseTypeName string
	FlagTemplate            string
	FieldDefinitions        string
}{
	{
		Name:                    "get",
		CallResolver:            true,
		IsConfig:                true,
		IsItr:                   false,
		IsSlice:                 false,
		APICallFuncName:         "GetTransferConfig",
		APICallResponseTypeName: "TransferConfig",
		Options: []string{
			"config-id",
			"name",
		},
	},
	{
		Name:                    "delete",
		CallResolver:            true,
		IsConfig:                true,
		IsItr:                   false,
		IsSlice:                 false,
		APICallFuncName:         "DeleteTransferConfig",
		APICallResponseTypeName: "TransferConfig",
		Options: []string{
			"config-id",
			"name",
		},
	},
	{
		Name:                    "history",
		CallResolver:            true,
		IsConfig:                false,
		IsItr:                   true,
		IsSlice:                 true,
		APICallFuncName:         "ListTransferRuns",
		APICallResponseTypeName: "TransferRun",
		Options: []string{
			"config-id",
			"name",
		},
	},
	{
		Name:                    "list",
		IsConfig:                true,
		IsItr:                   true,
		IsSlice:                 true,
		APICallFuncName:         "ListTransferConfigs",
		APICallResponseTypeName: "TransferConfig",
		Options:                 []string{},
	},
	{
		Name:                    "run",
		CallResolver:            true,
		IsConfig:                false,
		IsItr:                   false,
		IsSlice:                 false,
		APICallFuncName:         "StartManualTransferRuns",
		APICallResponseTypeName: "StartManualTransferRunsResponse",
		Options: []string{
			"config-id",
			"name",
			"since",
			"until",
			"at",
		},
	},
	{
		Name:                    "update",
		CallResolver:            true,
		IsConfig:                true,
		IsItr:                   false,
		IsSlice:                 false,
		APICallFuncName:         "UpdateTransferConfig",
		APICallResponseTypeName: "TransferConfig",
		Options: []string{
			"name",
			"config-id",
			"query",
			"dataset",
			"table",
			"partitioning-field",
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
	{
		Name:                    "create",
		IsConfig:                true,
		IsItr:                   false,
		IsSlice:                 false,
		APICallFuncName:         "CreateTransferConfig",
		APICallResponseTypeName: "TransferConfig",
		Options: []string{
			"name",
			"query",
			"dataset",
			"table",
			"partitioning-field",
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
