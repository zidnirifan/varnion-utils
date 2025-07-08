package influx

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func Ping(ctx context.Context, db influxdb2.Client) error {
	ready, err := db.Health(ctx)
	if err != nil {
		return err
	}

	if ready.Status != "pass" {
		return fmt.Errorf("influx not ready: %s", *ready.Message)
	}
	return nil
}
