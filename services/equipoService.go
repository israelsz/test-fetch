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
	CollectionNameEquipo = "Equipo"
)

// Función para crear un equipo e insertarlo a la base de datos de mongodb
func CreateEquipoService(newEquipo models.Equipo) (models.Equipo, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de equipos.
	collection := dbConnection.GetCollection(CollectionNameEquipo)

	// Genera un nuevo ID único para el equipo.
	newEquipo.ID = primitive.NewObjectID()
	// Establece la fecha de creación y actualización del equipo.
	newEquipo.CreatedAt = time.Now()
	newEquipo.UpdatedAt = time.Now()

	// Inserta el equipo en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newEquipo)
	//Si hubo un error
	if err != nil {
		return newEquipo, err
	}

	return newEquipo, err
}

// Función que devuelve el/los equipos a los que un evaluador puede evaluar
func GetEquiposAEvaluar(idEvaluador string) ([]models.Equipo, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameEquipo)
	// Variable que contiene a todos los equipo
	var equipo []models.Equipo
	// Se convierte el id a ObjectID
	oidEvaluador, err := primitive.ObjectIDFromHex(idEvaluador)
	if err != nil {
		return equipo, errors.New("no fue posible convertir el ID")
	}
	// Crea un filtro para buscar los equipos por id de equipo
	filter := bson.M{"idEvaluador": oidEvaluador}
	// Trae a todos los equipo desde la base de datos para el id de equipo entregado
	results, err := collection.Find(dbConnection.Context, filter)
	if err != nil {
		return equipo, errors.New("no fue posible traer a todos los equipo por el id de evaluador ingresado")
	}
	for results.Next(dbConnection.Context) {
		var singleEquipo models.Equipo
		if err = results.Decode(&singleEquipo); err != nil {
			log.Println("Equipo no se pudo añadir")
		}

		equipo = append(equipo, singleEquipo)
	}
	return equipo, nil
}

func GetAllEquiposCargosService() ([]models.ResponseEquipo, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameEquipo)
	// Variable que contiene a todos los equipos
	var responseEquipos []models.ResponseEquipo
	var singleResponseEquipo models.ResponseEquipo
	var cargo models.Cargo
	var singleResponseCargo models.ResponseCargo
	// Trae a todos los equipos desde la base de datos
	results, err := collection.Find(dbConnection.Context, bson.M{})
	if err != nil {
		return responseEquipos, errors.New("no fue posible traer a todos los equipos")
	}
	for results.Next(dbConnection.Context) {
		var singleEquipo models.Equipo
		if err = results.Decode(&singleEquipo); err != nil {
			log.Println("Equipo no se pudo añadir")
		}
		singleResponseEquipo.ID = singleEquipo.ID
		singleResponseEquipo.Name = singleEquipo.Name
		var responseCargos []models.ResponseCargo
		for _, id := range singleEquipo.Cargos {
			cargo, _ = GetCargoByIDService(id.Hex())
			singleResponseCargo.ID = cargo.ID
			singleResponseCargo.Name = cargo.Name
			responseCargos = append(responseCargos, singleResponseCargo)
		}
		singleResponseEquipo.Cargos = responseCargos
		responseEquipos = append(responseEquipos, singleResponseEquipo)
	}
	return responseEquipos, nil
}
