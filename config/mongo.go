package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Client {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(GetMongoEnv().MongoURI).
		SetMaxPoolSize(100). // Default is 100
		SetMinPoolSize(10).  // Default is 0
		SetMaxConnIdleTime(30*time.Minute))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
	return client
}

// Client instance
var MongoClient *mongo.Client = ConnectMongo()

// getting database collection by name
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(GetMongoEnv().Database).Collection(collectionName)
	return collection
}

func EnsureIndexes() {
	collection := GetCollection(MongoClient, "resume")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}
	log.Println("Index created successfully")
}
