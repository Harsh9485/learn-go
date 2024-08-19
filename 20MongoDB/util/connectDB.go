package util

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func ConnectDB(DatabaseName string, CollectionName string) *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:2017"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to Database seccessfully")
	Collection := client.Database(DatabaseName).Collection(CollectionName)
	return Collection
}

