package domain

import (
	"fmt"
	"regexp"
	"strings"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
)

const (
	parentFullFormat = `projects/%s/locations/%s`
	nameFullFormat   = `projects/%s/locations/%s/transferConfigs/%s`

	// TODO: avoid repetition
	nameFormatSuffix = `%s/transferConfigs/%s`

	nameFullFormatRegexp = `^projects/[^/]+/locations/[^/]+/transferConfigs/[^/]+$`
)

var nameFullFormatRegexpCompile = regexp.MustCompile(nameFullFormatRegexp)

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
	return fmt.Sprintf(nameFormatSuffix, p, c.ID), nil
}

func (c *TransferConfigsPathSpec) Parent() (string, error) {
	if c.ProjectID == "" || c.Location == "" {
		return "", fmt.Errorf("insufficient field values: projectID/ location")
	}
	return fmt.Sprintf(parentFullFormat, c.ProjectID, c.Location), nil
}

func GetTransferConfigIDByName(name string) (string, error) {
	if nameFullFormatRegexpCompile.MatchString(name) {
		return strings.Split(name, "/")[5], nil
	}
	return "", fmt.Errorf("name %s not match format %s", name, nameFullFormatRegexp)
}

func FindTransferConfigIDByName(tcList []*datatransferpb.TransferConfig, filterFunc func(*datatransferpb.TransferConfig) bool) []*datatransferpb.TransferConfig {
	var candidates []*datatransferpb.TransferConfig
	for _, m := range tcList {
		if filterFunc(m) {
			candidates = append(candidates, m)
		}
	}

	return candidates
}
