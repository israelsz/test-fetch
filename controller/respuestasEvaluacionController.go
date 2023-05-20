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
	var responseForm models.ResponseForm
	//Si ocurrio un error
	if err := ctx.ShouldBindJSON(&responseForm); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea el respuestasevaluacion en la base de datos Mongo
	createdResponseForm, err := services.CreateRespuestasEvaluacionService(responseForm)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el respuestasevaluacion"})
		return
	}
	log.Println("RespuestasEvaluacion creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdResponseForm)

}

// Servicio que permite recuperar un formulario para una competencia especifica
func GetFormularioEvaluacion(ctx *gin.Context) {
	// Se lee el id del cargo por el contexto ingresado en la ruta
	cargoID := ctx.Param("id")
	// Se crea el formulario a partir del cargo
	formularioEvaluacion, err := services.CreateQuestionsAnswersService(cargoID)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el formulario"})
		return
	}
	// Devuelve el formulario creado.
	ctx.JSON(http.StatusOK, formularioEvaluacion)
}
