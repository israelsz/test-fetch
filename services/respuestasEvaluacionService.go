package services

import (
	"rest-template/config"
	"rest-template/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Se establecen los nombres de la colección que se traeran desde la base de datos
*/
const (
	CollectionNameRespuestasEvaluacion = "RespuestasEvaluacion"
)

// Función para crear un gato e insertarlo a la base de datos de mongodb
func CreateRespuestasEvaluacionService(newRespuestasEvaluacion models.RespuestasEvaluacion) (models.RespuestasEvaluacion, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de gatos.
	collection := dbConnection.GetCollection(CollectionNameRespuestasEvaluacion)

	// Genera un nuevo ID único para el gato.
	newRespuestasEvaluacion.ID = primitive.NewObjectID()
	// Establece la fecha de creación y actualización del gato.
	newRespuestasEvaluacion.CreatedAt = time.Now()
	newRespuestasEvaluacion.UpdatedAt = time.Now()

	// Inserta el gato en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newRespuestasEvaluacion)
	//Si hubo un error
	if err != nil {
		return newRespuestasEvaluacion, err
	}

	return newRespuestasEvaluacion, err
}
