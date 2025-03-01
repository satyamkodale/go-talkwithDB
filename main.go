package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Ready to talk with DB ")

	// Define MongoDB connection URI
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.TODO())

	// err already exists, used = to update its value instead of :=
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Select the database and collection
	database := client.Database("talkwithDB")
	collection := database.Collection("user")

	// Insert a document
	newUser1 := bson.D{{"name", "Satyam"}, {"age", 24}, {"email", "satyam@example.com"}}
	insertResult1, err := collection.InsertOne(context.TODO(), newUser1)
	newUser2 := bson.D{{"name", "Satyam"}, {"age", 25}, {"email", "satyam38@example.com"}}
	insertResult2, err := collection.InsertOne(context.TODO(), newUser2)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a document with ID:", insertResult1.InsertedID)
	fmt.Println("Inserted a document with ID:", insertResult2.InsertedID)

	// bson.M It represents a flexible JSON-like document where
	// Find the document
	// Decode is a function that converts the MongoDB document into a Go variable.
	// bson.D (ordered key-value pairs) is used when order matters,
	var result bson.M
	err = collection.FindOne(context.TODO(), bson.D{{"name", "Satyam"}}).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found a document:", result)

}
