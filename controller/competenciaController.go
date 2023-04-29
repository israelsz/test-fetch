package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Servicio que permite Crear un competencia
func CreateCompetencia(ctx *gin.Context) {
	// Obtiene los datos del competencia a partir del cuerpo de la solicitud HTTP.
	var competencia models.Competencia
	//Si ocurrio un error
	if err := ctx.ShouldBindJSON(&competencia); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea el competencia en la base de datos Mongo
	createdCompetencia, err := services.CreateCompetenciaService(competencia)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el competencia"})
		return
	}
	log.Println("Competencia creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdCompetencia)

}
