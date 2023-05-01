package routes

import (
	"rest-template/controller"
	"rest-template/middleware"

	"github.com/gin-gonic/gin"
)

// InitRoutes registra las rutas junto a las funciones que ejecutan
func InitEquipoRoutes(r *gin.Engine) {
	// Define a group of routes with a shared set of middleware
	// Se define un grupo de rutas
	equipoGroup := r.Group("/equipo")
	{
		//Solo un usuario logueado sin importar su rol, puede crear un equipo
		equipoGroup.POST("/", middleware.LoadJWTAuth().MiddlewareFunc(), controller.CreateEquipo)
	}
}
