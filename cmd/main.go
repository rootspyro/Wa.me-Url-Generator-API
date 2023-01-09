package main

import (
	"wsurl-creator-api/handlers"
	"wsurl-creator-api/services"

	"github.com/gin-gonic/gin"
)

func main(){
	
	r := gin.Default()

	// Setup services
	wsServices := services.NewWsServices()
	
	r.GET("/ping", handlers.NewGetPingHandler().Handle)
	r.POST("/ws/url", handlers.NewGenUrlHandler(wsServices).Handle)

	r.Run()
}
