package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cargo struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Competencias []primitive.ObjectID      `json:"competencias" bson:"competencias,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
