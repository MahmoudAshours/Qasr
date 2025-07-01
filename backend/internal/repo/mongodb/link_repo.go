package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Link struct {
	Slug        string    `bson:"slug"`
	OriginalURL string    `bson:"original_url"`
	CreatedAt   time.Time `bson:"created_at"`
}

type LinkRepository struct {
	Collection *mongo.Collection
}

func NewLinkRepository(db *mongo.Database) *LinkRepository {
	return &LinkRepository{
		Collection: db.Collection("links"),
	}
}

func (r *LinkRepository) Create(slug, url string) error {
	_, err := r.Collection.InsertOne(context.TODO(), Link{
		Slug:        slug,
		OriginalURL: url,
		CreatedAt:   time.Now(),
	})
	return err
}

func (r *LinkRepository) FindBySlug(slug string) (string, error) {
	var link Link
	err := r.Collection.FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&link)
	if err != nil {
		return "", errors.New("not found")
	}
	return link.OriginalURL, nil
}
