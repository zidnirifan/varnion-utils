package elastic

import (
	"context"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

func Ping(ctx context.Context, db *elasticsearch.Client) error {
	ctxx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := db.Cluster.Health(
		db.Cluster.Health.WithContext(ctxx),
		db.Cluster.Health.WithPretty(),
	)
	if err != nil {
		return fmt.Errorf("elasticsearch ping failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("elasticsearch unhealthy: %s", res.Status())
	}

	return nil
}
