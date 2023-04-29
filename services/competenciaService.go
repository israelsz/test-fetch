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
	CollectionNameCompetencia = "Competencia"
)

// Función para crear un gato e insertarlo a la base de datos de mongodb
func CreateCompetenciaService(newCompetencia models.Competencia) (models.Competencia, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de gatos.
	collection := dbConnection.GetCollection(CollectionNameCompetencia)

	// Genera un nuevo ID único para el gato.
	newCompetencia.ID = primitive.NewObjectID()
	// Establece la fecha de creación y actualización del gato.
	newCompetencia.CreatedAt = time.Now()
	newCompetencia.UpdatedAt = time.Now()

	// Inserta el gato en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newCompetencia)
	//Si hubo un error
	if err != nil {
		return newCompetencia, err
	}

	return newCompetencia, err
}
