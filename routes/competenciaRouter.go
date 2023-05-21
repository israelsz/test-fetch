package routes

import (
	"rest-template/controller"
	"rest-template/middleware"

	"github.com/gin-gonic/gin"
)

// InitRoutes registra las rutas junto a las funciones que ejecutan
func InitCompetenciaRoutes(r *gin.Engine) {
	// Define a group of routes with a shared set of middleware
	// Se define un grupo de rutas
	competenciaGroup := r.Group("/competencia")
	{
		//Solo un usuario logueado sin importar su rol, puede crear un competencia
		competenciaGroup.POST("/", middleware.LoadJWTAuth().MiddlewareFunc(), controller.CreateCompetencia)

	}
}
