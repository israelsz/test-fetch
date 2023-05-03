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
	CollectionNameCargo = "Cargo"
)

// Función para crear un cargo e insertarlo a la base de datos de mongodb
func CreateCargoService(newCargo models.Cargo) (models.Cargo, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de cargos.
	collection := dbConnection.GetCollection(CollectionNameCargo)

	// Genera un nuevo ID único para el cargo.
	newCargo.ID = primitive.NewObjectID()
	// Establece la fecha de creación y actualización del cargo.
	newCargo.CreatedAt = time.Now()
	newCargo.UpdatedAt = time.Now()

	// Inserta el cargo en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newCargo)
	//Si hubo un error
	if err != nil {
		return newCargo, err
	}

	return newCargo, err
}

// Función para obtener un cargo por id
func GetCargoByIDService(cargoID string) (models.Cargo, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Crea un objeto ID de MongoDB a partir del ID del cargo.
	var cargo models.Cargo
	oid, err := primitive.ObjectIDFromHex(cargoID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return cargo, err
	}
	// Crea un filtro para buscar el cargo por su ID.
	filter := bson.M{"_id": oid}

	// Obtiene la colección de cargos.
	collection := dbConnection.GetCollection(CollectionNameCargo)
	err = collection.FindOne(dbConnection.Context, filter).Decode(&cargo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró ningún documento con el ID especificado.
			log.Println("Cargo no encontrado")
			return cargo, err
		}
		// Ocurrió un error durante la búsqueda.
		return cargo, err
	}
	log.Println("Se encontró el cargo")
	// Devuelve el cargo encontrado.
	return cargo, nil
}
