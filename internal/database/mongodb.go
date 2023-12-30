package database

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBLogOutput struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewMongoDBLogOutput(uri, databaseName, collectionName string) (*MongoDBLogOutput, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return &MongoDBLogOutput{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}, nil
}

func (m *MongoDBLogOutput) WriteDatabaseEntry(logEntry string) error {
	parts := strings.SplitN(logEntry, " ", 4)
	if len(parts) < 4 {
		return fmt.Errorf("log message format error")
	}

	timestamp, level, message := parts[0], parts[1], parts[3]
	logDocument := bson.M{
		"timestamp": timestamp,
		"level":     level,
		"message":   message,
	}

	collection := m.client.Database(m.databaseName).Collection(m.collectionName)
	_, err := collection.InsertOne(context.Background(), logDocument)
	return err
}

func (m *MongoDBLogOutput) Close() error {
	return m.client.Disconnect(context.Background())
}
