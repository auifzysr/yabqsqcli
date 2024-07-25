package domain

import (
	"encoding/json"
	"fmt"
)

type formatter interface {
	format(any) (string, error)
}

type Json struct{}

func (o *Json) format(d any) (string, error) {
	j, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

type PlainText struct{}

func (o *PlainText) format(d any) (string, error) {
	return fmt.Sprintf("%+v", d), nil
}

func Format(d any, formatRepr string) (string, error) {
	f, err := selectFormatter(formatRepr)
	if err != nil {
		return "", err
	}
	return f.format(d)
}

func selectFormatter(formatRepr string) (formatter, error) {
	switch formatRepr {
	case "plain":
		return &PlainText{}, nil
	case "json":
		return &Json{}, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", formatRepr)
	}
}
