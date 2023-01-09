package services

import (
	"fmt"
	"strings"
	"wsurl-creator-api/models"
)

type WSServices struct {}

func NewWsServices() *WSServices{
	return &WSServices{}
}

// this function parse the original message to an url format "Hello world" -> "Hello%20World"
func( ws *WSServices )FormatMessage(oMessage string) string{

	var formatedMessage string

	data := strings.Split(oMessage, " ")

	for i := range data {
		formatedMessage += data[i] + "%20"
	}

	return formatedMessage
}


func( ws *WSServices )CreateUrl(body *models.GenerateURLBody) string {

	formatedMessage := ws.FormatMessage(body.Message)

	url := fmt.Sprintf("https://wa.me/%s?text=%s", body.Phone, formatedMessage )

	return url
}
