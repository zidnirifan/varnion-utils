package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func Ping(ctx context.Context, db *mongo.Client) error {
	return db.Ping(ctx, nil)
}
