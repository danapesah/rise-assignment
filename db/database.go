package db

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

func GetDatabase() (m MongoDB) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://mongodb:27017").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

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

func (m MongoDB) Disconnect() {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	fmt.Println("Connection to DB disconnected!")
}

func (m MongoDB) Save(doc interface{}, collection string) {
	_, err := m.database.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
}

func (m MongoDB) Delete(structureInput interface{}, collection string) {
	filter, err := bson.Marshal(structureInput)

	_, err = m.database.Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
}

func (m MongoDB) Replace(structureInput interface{}, doc interface{}, collection string) {
	filter, err := bson.Marshal(structureInput)

	_, err = m.database.Collection(collection).ReplaceOne(context.TODO(), filter, doc)
	if err != nil {
		panic(err)
	}
}

func (m MongoDB) Load(structureInput interface{}, collection string) (resultCursor *mongo.Cursor, err error) {
	filter, err := bson.Marshal(structureInput)

	cursor, err := m.database.Collection(collection).Find(context.TODO(), filter)

	return cursor, err
}

func (m MongoDB) LoadByPagination(collection string, page int) (resultCursor *mongo.Cursor, err error) {
	filter := bson.D{}

	skip := page * 10
	opts := options.Find().SetLimit(int64(10)).SetSkip(int64(skip))
	cursor, err := m.database.Collection(collection).Find(context.TODO(), filter, opts)

	return cursor, err
}
