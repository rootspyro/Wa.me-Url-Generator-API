package services

import (
	"strings"
)

type CORSServices struct {
	
}

func NewCORSServices() *CORSServices {
	return &CORSServices{}
}

func( cs *CORSServices )ListAllowedOrigins( origins string ) []string {

	corsOrigins := strings.Split(origins, ",")

	return corsOrigins

}
