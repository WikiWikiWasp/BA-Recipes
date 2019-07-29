package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongo connects to a mongodb when given a URI
func ConnectToMongo() *mongo.Client {
	fmt.Println("Connecting to Mongo...")

	// set client options
	uri := "mongodb://localhost:27017/myrecipes"
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("...")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	return client
}

//CreateCollection creates a collection to hold the recipes
func CreateCollection(client *mongo.Client) *mongo.Collection {

	var db *mongo.Database
	coll := "recipes"
	dbs, err := client.ListDatabaseNames(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	} else {
		for _, db := range dbs {
			if db != "myrecipes" {
				continue
			}
		}
		fmt.Println("Database 'myrecipes' not found.")
		fmt.Println("Creating 'myrecipes' db...")
		db = client.Database("myrecipes")
	}

	// Get a handle for the `trainers` collection in the `test` db
	collection := client.Database(db.Name()).Collection(coll)

	return collection
}
