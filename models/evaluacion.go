package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Respuesta struct {
	Puntaje     int    `json:"puntaje" bson:"puntaje,omitempty"`
	Descripcion string `json:"descripcion" bson:"descripcion,omitempty"`
}

type Questions struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Tipo       string             `json:"tipo" bson:"tipo,omitempty"`
	Pregunta   string             `json:"pregunta" bson:"pregunta,omitempty"`
	Respuestas Respuesta          `json:"respuestas" bson:"respuestas,omitempty"`
}

type Evaluacion struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Tipo      string             `json:"tipo" bson:"tipo,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	Questions []Questions        `json:"questions" bson:"questions,omitempty"`
}
