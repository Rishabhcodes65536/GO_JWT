package routes

import (
	controllers "github.com/Rishabhcodes65536/GO_JWT/controllers"
	middleware "github.com/Rishabhcodes65536/GO_JWT/middleware"

	"github.com/gin-gonic/gin"
)
func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users",controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id",controllers.GetUser())
}

