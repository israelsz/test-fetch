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
	CollectionNameFormularioCompetencia = "FormularioCompetencia"
)

// Función para crear una formulariocompetencia e insertarla a la base de datos de mongodb
func CreateFormularioCompetenciaService(newFormularioCompetencia models.FormularioCompetencia) (models.FormularioCompetencia, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de formulariocompetencias.
	collection := dbConnection.GetCollection(CollectionNameFormularioCompetencia)

	// Genera un nuevo ID único para el formulariocompetencia.
	newFormularioCompetencia.ID = primitive.NewObjectID()
	// Establece la fecha de creación y actualización del formulariocompetencia.
	newFormularioCompetencia.CreatedAt = time.Now()
	newFormularioCompetencia.UpdatedAt = time.Now()

	// Inserta el formulariocompetencia en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newFormularioCompetencia)
	//Si hubo un error
	if err != nil {
		return newFormularioCompetencia, err
	}

	return newFormularioCompetencia, err
}
