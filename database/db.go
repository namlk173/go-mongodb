package database

import (
	"context"
	"go-mongodb/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Database struct {
	Client *mongo.Client
}

func NewDatabase() (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URL_DB))
	if err != nil {
		return &Database{}, err
	}
	return &Database{
		Client: client,
	}, nil
}

func (d *Database) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	defer func() {
		if err := d.Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
