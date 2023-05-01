package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Equipo struct {
	ID          primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name,omitempty"`
	IdEvaluador primitive.ObjectID   `json:"idEvaluador" bson:"idEvaluador,omitempty"`
	Cargos      []primitive.ObjectID `json:"cargos" bson:"cargos,omitempty"`
	CreatedAt   time.Time            `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time            `json:"updated_at" bson:"updated_at,omitempty"`
}
