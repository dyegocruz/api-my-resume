package resume

import (
	"context"

	"dyegocruz.com.br/api-my-resume/config"
	"go.mongodb.org/mongo-driver/bson"
)

func FindByUsername(username string) (MyResume, error) {
	collection := config.GetCollection(config.MongoClient, "resume")
	var resume MyResume
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&resume)
	if err != nil {
		return MyResume{}, err
	}
	return resume, nil
}
