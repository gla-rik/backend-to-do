package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glarik/backend-to-do/internal/routes"
)

func main() {
	// Инициализация Gin
	r := gin.Default()

	// Создание маршрутизатора
	router := routes.NewRouter()

	// Регистрация маршрутов
	router.RegisterRoutes(r)

	// Запуск сервера
	err := r.Run()
	if err != nil {
		return
	} // по умолчанию на порту 8080
}
