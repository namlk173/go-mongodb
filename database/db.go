package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Database struct {
	Client *mongo.Client
}

func NewDatabase(databaseURL string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	if err != nil {
		return &Database{}, err
	}

	return &Database{
		Client: client,
	}, nil
}

func (d *Database) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	defer func() {
		if err := d.Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func (d *Database) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := d.Client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	return nil
}
