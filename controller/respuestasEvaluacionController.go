package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Servicio que permite Crear un respuestasevaluacion
func CreateRespuestasEvaluacion(ctx *gin.Context) {
	// Obtiene los datos del respuestasevaluacion a partir del cuerpo de la solicitud HTTP.
	var respuestasevaluacion models.RespuestasEvaluacion
	//Si ocurrio un error
	if err := ctx.ShouldBindJSON(&respuestasevaluacion); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea el respuestasevaluacion en la base de datos Mongo
	createdRespuestasEvaluacion, err := services.CreateRespuestasEvaluacionService(respuestasevaluacion)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el respuestasevaluacion"})
		return
	}
	log.Println("RespuestasEvaluacion creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdRespuestasEvaluacion)

}
