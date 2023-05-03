package services

import (
	"errors"
	"log"
	"rest-template/config"
	"rest-template/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Se establecen los nombres de la colección que se traeran desde la base de datos
*/
const (
	CollectionNameCompetencia = "Competencia"
)

// Función para crear un competencia e insertarlo a la base de datos de mongodb
func CreateCompetenciaService(newCompetencia models.Competencia) (models.Competencia, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de competencias.
	collection := dbConnection.GetCollection(CollectionNameCompetencia)

	// Genera un nuevo ID único para el competencia.
	newCompetencia.ID = primitive.NewObjectID()
	// Establece la fecha de creación y actualización del competencia.
	newCompetencia.CreatedAt = time.Now()
	newCompetencia.UpdatedAt = time.Now()

	// Inserta el competencia en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newCompetencia)
	//Si hubo un error
	if err != nil {
		return newCompetencia, err
	}

	return newCompetencia, err
}

// Función para conseguir todas las competencias de un tipo (de tipo 1 - transversales | tipo 2- especificas)
func GetCompetenciasPorTipo(tipo int) ([]models.Competencia, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameCompetencia)
	// Variable que contiene a todos las competencias
	var competencias []models.Competencia
	// Crea un filtro para buscar la competencia por tipo
	filter := bson.M{"tipo": tipo}
	// Trae a todos los competencias desde la base de datos
	results, err := collection.Find(dbConnection.Context, filter)
	if err != nil {
		return competencias, errors.New("no fue posible traer a todas los competencias")
	}
	for results.Next(dbConnection.Context) {
		var singleCompetencia models.Competencia
		if err = results.Decode(&singleCompetencia); err != nil {
			log.Println("Competencia no se pudo añadir")
		}

		competencias = append(competencias, singleCompetencia)
	}
	return competencias, nil
}
