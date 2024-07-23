package domain

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/auifzysr/yabqsqcli/pkg/config"
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

func ResolveTransferConfigID(cfg config.Container) (string, error) {
	lcfg := &config.ListConfig{
		RootConfig: cfg.GetRootConfig(),
	}
	tcs := &TransferConfigsPathSpec{
		ProjectID: lcfg.ProjectID,
		Location:  lcfg.Region,
	}
	p, err := tcs.Parent()
	if err != nil {
		return "", err
	}

	tc := &datatransferpb.ListTransferConfigsRequest{
		Parent: p,
	}
	ctx := context.Background()

	var candidates []*datatransferpb.TransferConfig

	client, err := InitClient(ctx)
	if err != nil {
		return "", fmt.Errorf("data transfer client failed: %w", err)
	}
	itr := client.ListTransferConfigs(ctx, tc)
	for {
		m, err := itr.Next()
		if err != nil {
			break
		}
		if m.DisplayName == cfg.GetDisplayName() {
			candidates = append(candidates, m)
		}
	}
	switch len(candidates) {
	case 0:
		return "", fmt.Errorf("no such scheduled query: %s", cfg.GetDisplayName())
	case 1:
		return GetTransferConfigIDByName(candidates[0].Name)
	default:
		return "", fmt.Errorf("pick either of these: %+v", candidates)
	}
}
