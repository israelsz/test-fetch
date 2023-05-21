package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Servicio que permite Crear un formulariocompetencia
func CreateFormularioCompetencia(ctx *gin.Context) {
	// Obtiene los datos del formulariocompetencia a partir del cuerpo de la solicitud HTTP.
	var formulariocompetencia models.FormularioCompetencia
	//Si ocurrio un error
	if err := ctx.ShouldBindJSON(&formulariocompetencia); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea el formulariocompetencia en la base de datos Mongo
	createdFormularioCompetencia, err := services.CreateFormularioCompetenciaService(formulariocompetencia)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el formulariocompetencia"})
		return
	}
	log.Println("FormularioCompetencia creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdFormularioCompetencia)

}
