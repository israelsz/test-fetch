package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Servicio que permite Crear un equipo
func CreateEquipo(ctx *gin.Context) {
	// Obtiene los datos del equipo a partir del cuerpo de la solicitud HTTP.
	var equipo models.Equipo
	//Si ocurrio un error
	if err := ctx.ShouldBindJSON(&equipo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea el equipo en la base de datos Mongo
	createdEquipo, err := services.CreateEquipoService(equipo)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el equipo"})
		return
	}
	log.Println("Equipo creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdEquipo)

}

// Servicio para conseguir un listado de todos los equipos a los que un evaluador puede evaluar
func GetEquiposByEvaluadorID(ctx *gin.Context) {
	evaluadorID := ctx.Param("evaluadorid")
	equipos, err := services.GetEquiposAEvaluar(evaluadorID)
	// Si hubo un error
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró ningún documento con el ID especificado.
			log.Println("Equipos no encontrado")
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Equipos no encontrados"})
			return
		}
		// Ocurrió un error durante la búsqueda.
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Si no hay error se retorna a los usuarios
	log.Println("Se encontraron los Equipos")
	// Devuelve el usuario encontrado.
	ctx.JSON(http.StatusOK, equipos)
}

// Servicio para conseguir a todos los equipos de la base de datos
func GetAllEquiposCargos(ctx *gin.Context) {
	// Se consiguen a los equipos de la base de datos
	resultEquipos, err := services.GetAllEquiposCargosService()
	//Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los equipos"})
		return
	}
	//Se envia la respuesta http
	ctx.JSON(http.StatusCreated, resultEquipos)
}
