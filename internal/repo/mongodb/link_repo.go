package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Slug      string             `bson:"slug"`
	TargetURL string             `bson:"target_url"`
	Clicks    int                `bson:"clicks"`
	CreatedAt time.Time          `bson:"created_at"`
	ExpiresAt *time.Time         `bson:"expires_at,omitempty"`
}
