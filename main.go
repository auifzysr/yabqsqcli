package main

import (
	"context"
	"fmt"

	datatransfer "cloud.google.com/go/bigquery/datatransfer/apiv1"
	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
)

func main() {
	ctx := context.Background()
	c, err := datatransfer.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	m, err := c.GetTransferConfig(ctx, &datatransferpb.GetTransferConfigRequest{
		Name: fmt.Sprintf(`projects/%s/locations/%s/transferConfigs/%s`,
			projectID, region, transferConfigID),
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("meta: %+v", m)
}
