package config

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(cfg *Config) (*mongo.Client, error) {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(cfg.MongoDB.MongoURI).
		SetMaxPoolSize(100). // Default is 100
		SetMinPoolSize(10).  // Default is 0
		SetMaxConnIdleTime(30*time.Minute))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return client, nil
}

func EnsureIndexes(mongoClient *mongo.Client, cfg *Config) {
	collection := mongoClient.Database(cfg.MongoDB.Database).Collection("resume")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)

	if isIndexConflictError(err) {
		log.Println("Indexes already exist, skipping creation")
	} else {
		log.Println("Index created successfully")
	}
}

func isIndexConflictError(err error) bool {
	var cmdErr mongo.CommandError
	if errors.As(err, &cmdErr) {
		return cmdErr.Code == 85 || // IndexKeySpecsConflict
			cmdErr.Code == 86 || // IndexOptionsConflict
			cmdErr.Message == "index already exists"
	}
	return false
}
