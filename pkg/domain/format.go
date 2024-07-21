package domain

import (
	"encoding/json"
	"fmt"
)

type Formatter interface {
	Format(any) (string, error)
}

type Json struct{}

func (o *Json) Format(d any) (string, error) {
	j, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

type PlainText struct{}

func (o *PlainText) Format(d any) (string, error) {
	return fmt.Sprintf("%+v", d), nil
}

func SelectFormatter(format string) (Formatter, error) {
	switch format {
	case "plain":
		return &PlainText{}, nil
	case "json":
		return &Json{}, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}
