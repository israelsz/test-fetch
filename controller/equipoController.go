package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
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
