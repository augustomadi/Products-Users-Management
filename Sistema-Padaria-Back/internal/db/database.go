package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// ConnectDB establishes a connection to the MongoDB and returns an error if something goes wrong
func ConnectDB() error {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // substitua com a URI do seu MongoDB

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	return nil
}

// GetCollection returns a reference to a MongoDB collection
func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}

// IsDBConnected checks if the MongoDB connection is established
func IsDBConnected() bool {
	if client == nil {
		return false
	}

	// Check the connection
	err := client.Ping(context.TODO(), nil)
	return err == nil
}
