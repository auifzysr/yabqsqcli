package main

var configs = []struct {
	Name               string
	ClientCallTemplate string
	Options            []string
	FlagTemplate       string
}{
	{
		Name: "get",
		ClientCallTemplate: `
	m, err := client.GetTransferConfig(ctx, tc)
	if err != nil {
		return fmt.Errorf("get transfer failed: parent: %s, %w", fmt.Sprintf("projects/%s/locations/%s",
			cfg.ProjectID, cfg.Region,
		), err)
	}
	fmt.Printf("meta: %+v", m)
`,
		Options: []string{
			"config-id",
			"name",
		},
	},
	{
		Name: "delete",
		ClientCallTemplate: `
	err = client.DeleteTransferConfig(ctx, tc)
	if err != nil {
		return fmt.Errorf("deleting transfer failed: parent: %s, %w", fmt.Sprintf("projects/%s/locations/%s",
			cfg.ProjectID, cfg.Region,
		), err)
	}
`,
		Options: []string{
			"config-id",
			"name",
		},
	},
	{
		Name: "history",
		ClientCallTemplate: `
	itr := client.ListTransferRuns(ctx, tc)
	for {
		c, err := itr.Next()
		if err != nil {
			fmt.Printf("EOL or failed to iterate response: %s", err)
			break
		}
		fmt.Printf("run: %+v\n", c)
	}
`,
		Options: []string{
			"config-id",
			"name",
		},
	},
	{
		Name: "list",
		ClientCallTemplate: `
	itr := client.ListTransferConfigs(ctx, tc)
	for {
		c, err := itr.Next()
		if err != nil {
			fmt.Printf("EOL or failed to iterate response: %s", err)
			break
		}
		fmt.Printf("config: %+v\n", c)
	}
`,
		Options: []string{},
	},
	{
		Name: "run",
		ClientCallTemplate: `
	m, err := client.StartManualTransferRuns(ctx, tc)
	if err != nil {
		return fmt.Errorf("run transfer failed: parent: %s, %w", fmt.Sprintf("projects/%s/locations/%s",
			cfg.ProjectID, cfg.Region,
		), err)
	}
	fmt.Printf("meta: %+v", m)
`,
		Options: []string{
			"config-id",
			"name",
			"since",
			"until",
			"at",
		},
	},
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
		ClientCallTemplate: `
	m, err := client.CreateTransferConfig(ctx, tc)
	if err != nil {
		return fmt.Errorf("create transfer failed: parent: %s, %w", fmt.Sprintf("projects/%s/locations/%s",
			cfg.ProjectID, cfg.Region,
		), err)
	}
	fmt.Printf("meta: %+v", m)
`,
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
