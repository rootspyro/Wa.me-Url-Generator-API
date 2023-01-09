package models


type UrlData struct {
	ContactPhone string `json:"contact-phone"`
	Message string `json:"msg"`
	Url string `json:"url"`
}

type GeneratedURLResponse struct { 
	Status string `json:"status"`
	Data *UrlData	`json:"data"`
}
