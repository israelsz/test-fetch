package services

import (
	"log"
	"rest-template/config"
	"rest-template/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	// Genera un nuevo ID único para el formulario competencia.
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

// Función para conseguir el formulario de una competencia, buscando por id de competencia
func GetFormularioCompetenciaByCompetenciaIDService(idCompetencia primitive.ObjectID) (models.FormularioCompetencia, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Crea un objeto ID de MongoDB a partir del ID del formularioCompetencia.
	var formularioCompetencia models.FormularioCompetencia
	/*
		oid, err := primitive.ObjectIDFromHex(idCompetencia)
		if err != nil {
			log.Println("No fue posible convertir el ID")
			return formularioCompetencia, err
		}
	*/
	// Crea un filtro para buscar el formularioCompetencia por su ID.
	filter := bson.M{"idCompetencia": idCompetencia}

	// Obtiene la colección de formularioCompetencias.
	collection := dbConnection.GetCollection(CollectionNameFormularioCompetencia)
	err := collection.FindOne(dbConnection.Context, filter).Decode(&formularioCompetencia)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró ningún documento con el ID especificado.
			log.Println("formularioCompetencia no encontrado")
			return formularioCompetencia, err
		}
		// Ocurrió un error durante la búsqueda.
		return formularioCompetencia, err
	}
	log.Println("Se encontró el formularioCompetencia")
	// Devuelve el formularioCompetencia encontrado.
	return formularioCompetencia, nil
}

// Función para conseguir el formulario de una competencia, buscando por id de competencia
func GetFormularioCompetenciaByIDService(oid primitive.ObjectID) (models.FormularioCompetencia, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Crea un objeto ID de MongoDB a partir del ID del formularioCompetencia.
	var formularioCompetencia models.FormularioCompetencia
	/*
		oid, err := primitive.ObjectIDFromHex(idCompetencia)
		if err != nil {
			log.Println("No fue posible convertir el ID")
			return formularioCompetencia, err
		}
	*/
	// Crea un filtro para buscar el formularioCompetencia por su ID.
	filter := bson.M{"_id": oid}

	// Obtiene la colección de formularioCompetencias.
	collection := dbConnection.GetCollection(CollectionNameFormularioCompetencia)
	err := collection.FindOne(dbConnection.Context, filter).Decode(&formularioCompetencia)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró ningún documento con el ID especificado.
			log.Println("formularioCompetencia no encontrado")
			return formularioCompetencia, err
		}
		// Ocurrió un error durante la búsqueda.
		return formularioCompetencia, err
	}
	log.Println("Se encontró el formularioCompetencia")
	// Devuelve el formularioCompetencia encontrado.
	return formularioCompetencia, nil
}
