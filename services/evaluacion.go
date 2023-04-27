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
	CollectionNameEvaluacion = "Evaluacion"
)

// Función para crear una evaluacion e insertarla a la base de datos de mongodb
func CreateEvaluacionService(newEvaluacion models.Evaluacion) (models.Evaluacion, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de evaluacions.
	collection := dbConnection.GetCollection(CollectionNameEvaluacion)

	// Genera un nuevo ID único para el evaluacion.
	newEvaluacion.ID = primitive.NewObjectID()
	// Establece la fecha de creación y actualización del evaluacion.
	newEvaluacion.CreatedAt = time.Now()
	newEvaluacion.UpdatedAt = time.Now()

	// Inserta el evaluacion en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newEvaluacion)
	//Si hubo un error
	if err != nil {
		return newEvaluacion, err
	}

	return newEvaluacion, err
}
