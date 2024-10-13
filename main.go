/*
Модуль MusicInfo предоставляет API для получения информации о песнях.
Пользователи могут запрашивать данные о песнях, указывая название группы и название песни.
*/

// @title Music Info API
// @version 0.0.1
// @description API для получения информации о песнях по названию группы и песни. Была реализована в целях проверки корректности работы основной программы.
// @termsOfService http://swagger.io/terms/

// @contact.name Евгений
// @contact.url сайт пока отсутствует
// @contact.email i@evgesh4.ru

// @host localhost:9090
// @BasePath /
package main

import (
	_ "MusicInfo/docs" // Импортируйте ваши сгенерированные документы
	"MusicInfo/logger" // Импортируем логгер
	"MusicInfo/routes"
	"os"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Инициализируем логгер
	log := logger.InitLogger()

	// Загружаем переменные из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Получаем значение переменной API_PORT
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090" // Порт по умолчанию, если переменная окружения не установлена
	}

	// Настройка маршрутов с логгером
	router := routes.SetupRouter(log)

	// Регистрация Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера на указанном порту
	log.Infof("Starting server on port %s", port)
	router.Run(":" + port)
}
