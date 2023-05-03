package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RespuestasEvaluacion struct {
	ID                primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	IdEvaluador       primitive.ObjectID `json:"idEvaluador" bson:"idEvaluador,omitempty"`
	IdEvaluado        primitive.ObjectID `json:"idEvaluado" bson:"idEvaluado,omitempty"`
	TipoEvaluacion    string             `json:"tipoEvaluacion" bson:"tipoEvaluacion,omitempty"`
	Periodo           string             `json:"periodo" bson:"periodo,omitempty"`
	Retroalimentacion int                `json:"retroalimentacion" bson:"retroalimentacion,omitempty"`
	CreatedAt         time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt         time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	QuestionsAnswers  []struct {
		Competencia   string `json:"competencia" bson:"competencia,omitempty"`
		Respuesta     string `json:"respuesta" bson:"respuesta,omitempty"`
		Puntaje       int    `json:"puntaje" bson:"puntaje,omitempty"`
		Justificacion string `json:"justificacion" bson:"justificacion,omitempty"`
	} `json:"questionsAnswers" bson:"questionsAnswers,omitempty"`
	Evaluacion FormularioCompetencia `json:"evaluacion" bson:"evaluacion,omitempty"`
}
