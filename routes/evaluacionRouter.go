package routes

import (
	"rest-template/controller"
	"rest-template/middleware"

	"github.com/gin-gonic/gin"
)

// InitRoutes registra las rutas junto a las funciones que ejecutan
func InitEvaluacionRoutes(r *gin.Engine) {
	// Define a group of routes with a shared set of middleware
	// Se define un grupo de rutas
	evaluacionGroup := r.Group("/evaluacion")
	{
		//Solo un usuario logueado sin importar su rol, puede crear un evaluacion
		evaluacionGroup.POST("/", middleware.LoadJWTAuth().MiddlewareFunc(), controller.CreateEvaluacion)

	}
}
