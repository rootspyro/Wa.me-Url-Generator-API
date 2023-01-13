package services

import (
	"fmt"
	"regexp"
	"strings"
	"wsurl-creator-api/models"
)

type WSServices struct {}

func NewWsServices() *WSServices{
	return &WSServices{}
}
 
func( ws *WSServices )ValidateNumberFormat(phoneNumber string) bool {

	/*
		This function validates the correct format of phone number
		
		bad: +1 (xxx)-xxx-xxxx 
		bad: +1 xxx-xxx-xxxx 
		bad: 1 xxx-xxx-xxxx 

		good: +1xxxxxxxxxx
		good: 1xxxxxxxxxx

	*/

	
	validator := regexp.MustCompile(`^\+(?:[0-9]\x20?){6,14}[0-9]$`)
	
	return validator.MatchString(phoneNumber)
}

// this function parse the original message to an url format "Hello world" -> "Hello%20World"
func( ws *WSServices )FormatMessage(oMessage string) string{

	var formatedMessage string

	data := strings.Split(oMessage, " ")

	for i := range data {
		formatedMessage += data[i] 
		if i < len(data) - 1 {
			formatedMessage += "%20"
		}
	}

	return formatedMessage
}


func( ws *WSServices )CreateUrl(body *models.GenerateURLBody) string {

	formatedMessage := ws.FormatMessage(body.Message)

	url := fmt.Sprintf("https://wa.me/%s?text=%s", body.Phone, formatedMessage )

	return url
}


