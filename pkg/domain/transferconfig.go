package domain

import "fmt"

type TransferConfigsPathSpec struct {
	ProjectID string
	Location  string
	ID        string
}

func (c *TransferConfigsPathSpec) Name() (string, error) {
	if c.ID == "" {
		return "", fmt.Errorf("insufficient field values: transferConfigID")
	}
	p, err := c.Parent()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`%s/transferConfigs/%s`, p, c.ID), nil
}

func (c *TransferConfigsPathSpec) Parent() (string, error) {
	if c.ProjectID == "" || c.Location == "" {
		return "", fmt.Errorf("insufficient field values: projectID/ location")
	}
	return fmt.Sprintf(`projects/%s/locations/%s`,
		c.ProjectID, c.Location), nil
}
