package routes

import (
	"MusicInfo/controller"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// SetupRouter - регистрирует маршруты API для приложения MusicInfo.
// Этот метод добавляет маршруты к переданному объекту маршрутизатора.
func SetupRouter(logger *logrus.Logger) *gin.Engine {
	r := gin.Default()
	logger.Info("Setting up API routes")
	r.GET("/info", controller.GetSongInfo(logger)) // Регистрирует маршрут для получения информации о песне
	return r
}
