package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishansaini194/Projects/controllers"
	"github.com/ishansaini194/Projects/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
