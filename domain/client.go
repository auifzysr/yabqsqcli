package domain

import (
	"context"

	datatransfer "cloud.google.com/go/bigquery/datatransfer/apiv1"
)

type Client struct {
	*datatransfer.Client
}

func InitClient(ctx context.Context) (*Client, error) {
	c, err := datatransfer.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{c}, nil
}

func (c *Client) Close() {
	c.Client.Close()
}
