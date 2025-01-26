package router

import (
	"github.com/gin-gonic/gin"
	"github.com/glarik/backend-to-do/internal/controllers"
)

type Router struct {
	ginInstance *gin.Engine
}

// Конструктор для Router
func NewRouter(ginInstance *gin.Engine, pingController *controller.PingController) *Router {
	router := &Router{
		ginInstance: ginInstance,
	}
	router.InitRoutes(pingController)
	return router
}

// Инициализация маршрутов
func (router *Router) InitRoutes(pingController *controller.PingController) {
	router.ginInstance.GET("/ping", pingController.Ping)
}
