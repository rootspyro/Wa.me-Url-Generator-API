package main

import (
	"os"
	"wsurl-creator-api/handlers"
	"wsurl-creator-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	
	r := gin.Default()

	// Load .env file
	godotenv.Load()

	// Setup services
	wsServices := services.NewWsServices()
	corsServices := services.NewCORSServices()

	corsOrigins := corsServices.ListAllowedOrigins(os.Getenv("ALLOWED_ORIGINS"))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigins,
    AllowMethods:     []string{"GET","POST"},
    AllowHeaders:     []string{"Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "https://github.com"
    },
	}))

	// Setup Handlers

	v1 := r.Group("/api/v1") 

	v1.GET("/ping", handlers.NewGetPingHandler().Handle)
	v1.POST("/ws/url", handlers.NewGenUrlHandler(wsServices).Handle)


	// Setup CORS Policy

	r.Run()
}
