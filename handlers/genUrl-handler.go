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
			Data: err.Error(),
		}

		c.JSON(500, errResponse)

		return
	}

	if !h.service.ValidateNumberFormat(body.Phone) { 

		errResponse := &models.Default{
			Status: "error",
			Data: "Invalid phone number format! Avoid spaces and special characters, except for the `+` character in the international code.",
		}

		c.JSON(500, errResponse)
		return
	}

	if len(body.Message) > 250 {
		
		errResponse := &models.Default{
			Status: "error",
			Data: "Message length over 250 characters!",
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
