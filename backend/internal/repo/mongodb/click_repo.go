package mongodb

import (
	"context"
	model "qasr/backend/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClickRepo struct {
	Collection *mongo.Collection
}

func NewClickRepo(db *mongo.Database) *ClickRepo {
	return &ClickRepo{
		Collection: db.Collection("clicks"),
	}
}

func (r *ClickRepo) SaveClick(click *model.Click) error {
	_, err := r.Collection.InsertOne(context.TODO(), click)
	return err
}

func (r *ClickRepo) GetClicksBySlug(slug string) ([]*model.Click, error) {
	cursor, err := r.Collection.Find(context.TODO(), bson.M{"slug": slug})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var clicks []*model.Click
	if err := cursor.All(context.TODO(), &clicks); err != nil {
		return nil, err
	}
	return clicks, nil
}
