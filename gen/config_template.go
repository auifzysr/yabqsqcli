package main

var configTemplate = `
// Code generated by gen/main.go; DO NOT EDIT.

package config

type {{ .Name | capitalize }}Config struct {
	*RootConfig

	{{ .FieldDefinitions }}
}
`
