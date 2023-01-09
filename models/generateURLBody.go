package models

type GenerateURLBody struct {
	Phone string `json:"phone" binding:"required"`
	Message string `json:"msg" binding:"required"`
}
