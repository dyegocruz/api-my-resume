package repositories

import (
	"context"

	"dyegocruz.com.br/api-my-resume/internal/config"
	"dyegocruz.com.br/api-my-resume/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type resumeRepo struct {
	db  *mongo.Client // or any other DB client
	cfg *config.Config
}

func NewResumeRepository(db *mongo.Client, cfg *config.Config) MyResumeRepository {
	return &resumeRepo{db: db, cfg: cfg}
}

func (r *resumeRepo) FindByUsername(username string) (models.MyResume, error) {
	collection := r.db.Database(r.cfg.MongoDB.Database).Collection("resume")
	var resume models.MyResume
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&resume)
	if err != nil {
		return models.MyResume{}, err
	}
	return resume, nil
}
