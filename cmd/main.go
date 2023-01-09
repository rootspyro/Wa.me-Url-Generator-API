package main

import (
	"os"
	"wsurl-creator-api/handlers"
	"wsurl-creator-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	
	r := gin.Default()

	// Load env variables from .env
	godotenv.Load()

	// Setup services
	wsServices := services.NewWsServices()
	
	r.GET("/ping", handlers.NewGetPingHandler().Handle)
	r.POST("/ws/url", handlers.NewGenUrlHandler(wsServices).Handle)

	r.Run(":" + os.Getenv("PORT"))
}
