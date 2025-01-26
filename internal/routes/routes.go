package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/glarik/backend-to-do/internal/controllers"
)

type Router struct{}

func NewRouter() *Router {
	router := &Router{}
	return router
}

func (r *Router) RegisterRoutes(router *gin.Engine) {
	// Создаем экземпляр PingController
	pingController := controllers.NewPingController()

	// Привязываем маршрут "/ping" к методу Ping в PingController
	router.GET("/ping", pingController.Ping)
}
