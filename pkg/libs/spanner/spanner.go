package spanner

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"github.com/yudai2929/grpc-practice/pkg/config"
)

func NewClient() (*spanner.Client, error) {
	ctx := context.Background()

	projectID := config.Env.Spanner.ProjectID
	instanceName := config.Env.Spanner.InstanceName
	dbName := config.Env.Spanner.DBName

	fullDBName := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectID, instanceName, dbName)

	client, err := spanner.NewClient(ctx, fullDBName)
	if err != nil {
		return nil, err
	}

	return client, nil
}
