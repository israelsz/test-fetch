package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Servicio que permite Crear un evaluacion
func CreateEvaluacion(ctx *gin.Context) {
	// Obtiene los datos del evaluacion a partir del cuerpo de la solicitud HTTP.
	var evaluacion models.Evaluacion
	//Si ocurrio un error
	if err := ctx.ShouldBindJSON(&evaluacion); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea el evaluacion en la base de datos Mongo
	createdEvaluacion, err := services.CreateEvaluacionService(evaluacion)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el evaluacion"})
		return
	}
	log.Println("Evaluacion creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdEvaluacion)

}
