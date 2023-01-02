package repo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
	db     *mongo.Database
}

func New(dbUri string, dbName string) (*Repository, error) {
	log.Println("connecting db")
	clientOptions := options.Client()
	clientOptions.ApplyURI(dbUri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return &Repository{
		client: client,
		db:     client.Database(dbName),
	}, nil
}
