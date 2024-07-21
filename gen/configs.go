package main

var scenarios = []struct {
	Name               string
	ClientCallTemplate string
	Options            []string
	FlagTemplate       string
	FieldDefinitions   string
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
	f, err := domain.SelectFormatter(cfg.OutputFormat)
	if err != nil {
		return err
	}
	o, err := f.Format(m)
	if err != nil {
		return err
	}
	fmt.Printf("%s", o)
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
		m, err := itr.Next()
		if err != nil {
			fmt.Printf("EOL or failed to iterate response: %s", err)
			break
		}
		f, err := domain.SelectFormatter(cfg.OutputFormat)
		if err != nil {
			return err
		}
		o, err := f.Format(m)
		if err != nil {
			return err
		}
		fmt.Printf("%s", o)
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
		m, err := itr.Next()
		if err != nil {
			fmt.Printf("EOL or failed to iterate response: %s", err)
			break
		}
		f, err := domain.SelectFormatter(cfg.OutputFormat)
		if err != nil {
			return err
		}
		o, err := f.Format(m)
		if err != nil {
			return err
		}
		fmt.Printf("%s", o)
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
	f, err := domain.SelectFormatter(cfg.OutputFormat)
	if err != nil {
		return err
	}
	o, err := f.Format(m)
	if err != nil {
		return err
	}
	fmt.Printf("%s", o)
`,
		Options: []string{
			"config-id",
			"name",
			"since",
			"until",
			"at",
		},
	},
	{
		Name: "update",
		ClientCallTemplate: `
	m, err := client.UpdateTransferConfig(ctx, tc)
	if err != nil {
		return fmt.Errorf("updating transfer failed: name=%s, err=%w",
			fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s",
				cfg.ProjectID, cfg.Region, cfg.TransferConfigID), err)
	}
	f, err := domain.SelectFormatter(cfg.OutputFormat)
	if err != nil {
		return err
	}
	o, err := f.Format(m)
	if err != nil {
		return err
	}
	fmt.Printf("%s", o)
`,
		Options: []string{
			"name",
			"config-id",
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
	{
		Name: "create",
		ClientCallTemplate: `
	m, err := client.CreateTransferConfig(ctx, tc)
	if err != nil {
		return fmt.Errorf("create transfer failed: parent: %s, %w", fmt.Sprintf("projects/%s/locations/%s",
			cfg.ProjectID, cfg.Region,
		), err)
	}
	f, err := domain.SelectFormatter(cfg.OutputFormat)
	if err != nil {
		return err
	}
	o, err := f.Format(m)
	if err != nil {
		return err
	}
	fmt.Printf("%s", o)
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
