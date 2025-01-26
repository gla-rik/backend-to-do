package controllers

import (
	"github.com/gin-gonic/gin"
)

// PingController структура контроллера для обработки запросов ping
type PingController struct{}

// NewPingController Конструктор для PingController
func NewPingController() *PingController {
	return &PingController{}
}

// Ping Метод для обработки GET-запросов
func (pingController *PingController) Ping(c *gin.Context) {
	// Логика ответа
	c.JSON(200, gin.H{
		"message": "pong", // JSON-объект, который возвращается в ответе
	})
}
