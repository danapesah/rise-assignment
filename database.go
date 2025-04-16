package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDB struct {
	client   *mongo.Client
	database *mongo.Database
}

func getDatabase() (m MongoDB) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	instance := MongoDB{
		client:   client,
		database: client.Database("admin"),
	}

	return instance
}

func (m MongoDB) disconnect() {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	fmt.Println("Connection to DB disconnected!")
}

func (m MongoDB) save(doc interface{}, collection string) {
	_, err := m.database.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
}

func (m MongoDB) delete(doc interface{}, collection string) {
	_, err := m.database.Collection(collection).DeleteOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
}

func (m MongoDB) replace(doc interface{}, collection string) {
	_, err := m.database.Collection(collection).ReplaceOne(context.TODO(), doc, doc)
	if err != nil {
		panic(err)
	}
}

func (m MongoDB) load(id int, collection string) () {
	if id != 0 {
		filter := bson.D{{"_id", id}}
	} else {
		filter := bson.D{{"$limit", 3}}
	}

	var contact Contact
	err := m.database.Collection(collection).FindOne(context.TODO(), filter).Decode(&contact)
	if err != nil {
		panic(err)
	}

	fmt.Println(contact)
}
