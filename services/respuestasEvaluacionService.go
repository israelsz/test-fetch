package services

import (
	"errors"
	"log"
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
func CreateRespuestasEvaluacionService(newResponseForm models.ResponseForm) (models.RespuestasEvaluacion, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de gatos.
	collection := dbConnection.GetCollection(CollectionNameRespuestasEvaluacion)

	// Objeto que se insertara en la base de datos
	var newRespuestasEvaluacion models.RespuestasEvaluacion
	// Genera un nuevo ID único para el gato.
	newRespuestasEvaluacion.ID = primitive.NewObjectID()
	// Se setean los valores del response
	newRespuestasEvaluacion.IdEvaluado = newResponseForm.IdEvaluado
	newRespuestasEvaluacion.IdEvaluador = newResponseForm.IdEvaluador
	newRespuestasEvaluacion.TipoEvaluacion = newResponseForm.TipoEvaluacion
	newRespuestasEvaluacion.Periodo = newResponseForm.Periodo
	newRespuestasEvaluacion.Retroalimentacion = newResponseForm.Retroalimentacion
	newRespuestasEvaluacion.QuestionsAnswers = newResponseForm.QuestionsAnswers
	// Establece la fecha de creación y actualización del gato.
	newRespuestasEvaluacion.CreatedAt = time.Now()
	newRespuestasEvaluacion.UpdatedAt = time.Now()
	// Se recupera los formularios
	var formulariosCompetencia []models.FormularioCompetencia
	var formulario models.FormularioCompetencia
	log.Println("IDS:", newResponseForm.IdsFormularios)
	for _, id := range newResponseForm.IdsFormularios {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println("No fue posible convertir el ID")
			return newRespuestasEvaluacion, errors.New("id invalido")
		}
		formulario, _ = GetFormularioCompetenciaByIDService(oid)
		log.Println("Id:", id, ",FORMULARIOS:", formulario)
		formulariosCompetencia = append(formulariosCompetencia, formulario)
	}
	// Se setean los formularios
	newRespuestasEvaluacion.Evaluacion = formulariosCompetencia

	// Inserta el gato en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newRespuestasEvaluacion)
	//Si hubo un error
	if err != nil {
		return newRespuestasEvaluacion, err
	}

	return newRespuestasEvaluacion, err
}

// Función para crear formulario para enviar al frontend
func CreateQuestionsAnswersService(cargoID string) (models.ResponseForm, error) {
	// Se inicializa modelo de respuesta a enviar
	var formResponse models.ResponseForm
	// Se trae el objeto cargo
	cargo, err := GetCargoByIDService(cargoID)
	//Si hubo un error o el cargo no se encuentra
	if err != nil {
		return formResponse, err
	}
	idsCompetencias := cargo.Competencias
	var tipoCompetencias []string
	// Agrega el tipo de competencia cargada
	for i := 0; i < len(idsCompetencias); i++ {
		tipoCompetencias = append(tipoCompetencias, "Especifica")
	}
	// Agregar id's de competencias transversales
	competenciasTransversales, err := GetCompetenciasPorTipo(1) // 1 = Competencias transversales
	//Si hubo un error
	if err != nil {
		return formResponse, err
	}
	// Se agregan los ids de las compentencias transversales a la lista de ids
	for _, competenciaTransversal := range competenciasTransversales {
		idsCompetencias = append(idsCompetencias, competenciaTransversal.ID)
		tipoCompetencias = append(tipoCompetencias, "Transversal")
	}

	// Agregar id's de competencias de mejora
	competenciasMejora, err := GetCompetenciasPorTipo(3) // 3 = Competencias oportunidades de mejora
	//Si hubo un error
	if err != nil {
		return formResponse, err
	}
	// Se agregan los ids de las compentencias oportunidades de mejora a la lista de ids
	for _, competenciaMejora := range competenciasMejora {
		idsCompetencias = append(idsCompetencias, competenciaMejora.ID)
		tipoCompetencias = append(tipoCompetencias, "Mejora")
	}
	// Variable que contiene a todos los formulariosCompetencia a recolectar
	var formularioObtenido models.FormularioCompetencia
	var idsFormularios []string
	// Crear preguntas - respuestas para enviar (QuestionAnswers)
	var preguntasRespuestas []models.QuestionsAnswers
	var preguntaRepuesta models.QuestionsAnswers
	// Se consiguen todos los formularios para el id de competencia
	for _, id := range idsCompetencias {
		formularioObtenido, _ = GetFormularioCompetenciaByCompetenciaIDService(id)
		idsFormularios = append(idsFormularios, formularioObtenido.ID.Hex())
		preguntaRepuesta.Competencia = formularioObtenido.Questions[0].Pregunta //Texto Pregunta
		preguntaRepuesta.Justificacion = "-1"
		preguntaRepuesta.Puntaje = -1
		preguntaRepuesta.OpcionesPregunta = formularioObtenido.Questions[0].Respuestas //Opciones que puede tomar la pregunta
		preguntasRespuestas = append(preguntasRespuestas, preguntaRepuesta)
	}
	formResponse.QuestionsAnswers = preguntasRespuestas
	formResponse.IdsFormularios = idsFormularios
	formResponse.TipoCompetencia = tipoCompetencias
	return formResponse, nil
}
