package controller

import "github.com/gin-gonic/gin"

type PingController struct{}

// Конструктор для PingController
func NewPingController() *PingController {
	return &PingController{}
}

// Обработчик для эндпоинта /ping
func (pingController *PingController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
