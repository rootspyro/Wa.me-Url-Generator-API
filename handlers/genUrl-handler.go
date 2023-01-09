package handlers

import (
	"log"
	"wsurl-creator-api/models"
	"wsurl-creator-api/services"

	"github.com/gin-gonic/gin"
)

type GenUrlHandler struct{
	service *services.WSServices
}

func NewGenUrlHandler(s *services.WSServices) *GenUrlHandler{
	return &GenUrlHandler{
		service: s,
	}
}

func( h *GenUrlHandler )Handle(c *gin.Context) {

	var body *models.GenerateURLBody

	if err := c.BindJSON(&body); err != nil {

		log.Println(err)
		
		errResponse := &models.Default{
			Status: "error",
			Data: "Body error",
		}

		c.JSON(500, errResponse)

		return
	}

	url := h.service.CreateUrl(body) 
	
	data := &models.UrlData{
		ContactPhone: body.Phone,
		Message: body.Message,
		Url: url,
	} 
	
	oKResponse := &models.GeneratedURLResponse{
		Status: "success",
		Data: data,
	}

	c.JSON(200, oKResponse)
}
