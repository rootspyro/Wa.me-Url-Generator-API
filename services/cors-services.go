package services

import (
	"log"
	"strings"
)

type CORSServices struct {
	
}

func NewCORSServices() *CORSServices {
	return &CORSServices{}
}

func( cs *CORSServices )ListAllowedOrigins( origins string ) []string {

	corsOrigins := strings.Split(origins, ",")
	log.Println(corsOrigins)
	return corsOrigins

}
