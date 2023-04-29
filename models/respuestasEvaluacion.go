package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RespuestasEvaluacion struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Evaluacion string             `json:"evaluacion" bson:"evaluacion,omitempty"`
	Periodo    string             `json:"periodo" bson:"periodo,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	User       User               `json:"user" bson:"user,omitempty"`
	Preguntas  []struct {
		Pregunta  string `json:"pregunta" bson:"pregunta"`
		Respuesta string `json:"respuesta" bson:"respuesta"`
	} `json:"preguntas" bson:"preguntas"`
}
