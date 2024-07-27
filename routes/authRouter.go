package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/Rishabhcodes65536/GO_JWT/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup",controllers.SignUp())
	incomingRoutes.POST("users/login",controllers.Login())
}