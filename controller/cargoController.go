package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Servicio que permite Crear un cargo
func CreateCargo(ctx *gin.Context) {
	// Obtiene los datos del cargo a partir del cuerpo de la solicitud HTTP.
	var cargo models.Cargo
	//Si ocurrio un error
	if err := ctx.ShouldBindJSON(&cargo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea el cargo en la base de datos Mongo
	createdCargo, err := services.CreateCargoService(cargo)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el cargo"})
		return
	}
	log.Println("Cargo creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdCargo)

}
