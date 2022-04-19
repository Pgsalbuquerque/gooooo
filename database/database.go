package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func StartDB() {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://pgsa:9932@localhost:27017"))

	if err != nil {
		fmt.Println("Could not connect to the Mongo")
		log.Fatal("Error: ", err)
	}

	db = client
}

func GetDB() *mongo.Client {
	return db
}
