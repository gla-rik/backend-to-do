package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/glarik/backend-to-do/internal/controllers"
	"github.com/glarik/backend-to-do/internal/router"
)

func main() {
	// Инициализация gin
	ginInstance := gin.Default()

	// Создание экземпляра PingController
	pingController := controller.NewPingController()

	// Создание экземпляра Router
	_ = router.NewRouter(ginInstance, pingController)

	// Запуск сервера
	ginInstance.Run()
}
