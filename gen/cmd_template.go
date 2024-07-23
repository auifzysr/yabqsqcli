package main

var cmdTemplate = `
// Code generated by gen/main.go; DO NOT EDIT.

package cmd

import (
	"context"
	"fmt"

	"github.com/auifzysr/yabqsqcli/pkg/config"
	"github.com/auifzysr/yabqsqcli/pkg/factory"
	"github.com/auifzysr/yabqsqcli/pkg/domain"
    {{- if or (eq .Name "list") (eq .Name "history")}}
	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	{{- end }}
	"github.com/urfave/cli/v2"
)

func {{ .Name }}(cfg *config.{{ .Name | capitalize }}Config) error {
    {{- if .CallResolver}}
	if cfg.TransferConfigID == "" {
		configID, err := domain.ResolveTransferConfigID(cfg)
		if err != nil {
			return err
		}
		cfg.TransferConfigID = configID
	}
	{{- end}}
	tc, err := factory.{{ .Name | capitalize }}TransferConfigFactory(cfg)
	if err != nil {
		return err
	}
	ctx := context.Background()

	{{ .ClientCallTemplate }}

	return nil
}

func {{ .Name }}Command(rootCfg *config.RootConfig) *cli.Command {
	cfg := &config.{{ .Name | capitalize }}Config{
		RootConfig: rootCfg,
	}

	return &cli.Command{
		Name:    "{{ .Name }}",
		Description:   "{{ .Name }} scheduled query config",
		Action: func(cCtx *cli.Context) error {
			return {{ .Name }}(cfg)
		},
		Flags: []cli.Flag{
		{{ .FlagTemplate }}
		},
	}
}
`
