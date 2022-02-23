package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDBHandler struct {
	dbClient     *mongo.Client
	dbCollection *mongo.Collection
}

func newMongoDBHandler() (*mongoDBHandler, error) {
	fmt.Println("Connecting to MongoDB")

	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	if err := dbClient.Connect(context.Background()); err != nil {
		return nil, err
	}

	dbCollection := dbClient.Database("mydb").Collection("todo")

	return &mongoDBHandler{
		dbClient,
		dbCollection,
	}, nil
}

func (h *mongoDBHandler) disconnect() {
	if err := h.dbClient.Disconnect(context.Background()); err != nil {
		log.Fatalf("faied to disconnect MongoDB %v", err)
	}
}
