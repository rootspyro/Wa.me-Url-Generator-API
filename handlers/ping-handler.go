package handlers

import "github.com/gin-gonic/gin"

type GetPingHandler struct {

}

func NewGetPingHandler() *GetPingHandler {
	return &GetPingHandler{}
}

func( h *GetPingHandler ) Handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
		"data": "pong", 
	})
} 
