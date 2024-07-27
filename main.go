package main

import (
	"log"
	"os"

	routes "github.com/Rishabhcodes65536/GO_JWT/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(".env")
	if err!=nil {
		log.Fatal("Error loading .env file")
	}
	PORT:=os.Getenv("PORT")

	if PORT==""{
		PORT="8000"
	}
	
	router:=gin.New()
	router.Use(gin.Logger())


	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func (c *gin.Context)  {
		c.JSON(200, gin.H{"success":"Access granted for API-1"})
	})

	router.GET("/api-2", func (c *gin.Context)  {
		c.JSON(200, gin.H{"success":"Access granted for API-2"})
	})

	router.Run(":"+ PORT)

}