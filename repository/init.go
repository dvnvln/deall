package repo

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
	dbName string
}

func New(dbUri string, dbName string) (*Repository, error) {
	log.Println("connecting db")
	clientOptions := options.Client()
	clientOptions.ApplyURI(dbUri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	return &Repository{
		client: client,
		dbName: dbName,
	}, nil
}
