package controller

import (
	"MusicInfo/database"
	"MusicInfo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetSongInfo - обработчик для получения информации о песне.
// Этот метод извлекает параметры запроса "group" и "song",
// проверяет их наличие и выполняет поиск песни в базе данных.
//
// @Summary Получить информацию о песне
// @Description Возвращает информацию о песне по группе и названию.
// @Tags songs
// @Accept json
// @Produce json
// @Param group query string true "Название группы"
// @Param song query string true "Название песни"
// @Success 200 {object} models.SongDetail "Информация о песне"
// @Failure 400 {object} models.ErrorResponse "Ошибка: отсутствуют необходимые параметры"
// @Failure 404 {object} models.ErrorResponse "Ошибка: песня не найдена"
// @Router /info [get]
func GetSongInfo(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		group := c.Query("group") // Извлечение параметра "group" из запроса
		song := c.Query("song")   // Извлечение параметра "song" из запроса

		// Проверка на наличие параметров
		if group == "" || song == "" {
			logger.Warn("Missing required parameters: group and song")
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Missing required query parameters: group and song"})
			return // Возвращаем ошибку 400, если параметры отсутствуют
		}

		// Поиск песни в "базе данных"
		logger.Debugf("Searching for song: group=%s, song=%s", group, song)
		if details, exists := database.GetSongDetail(group, song); exists {
			logger.Infof("Song found: %s - %s", group, song)
			c.JSON(http.StatusOK, details) // Возвращаем информацию о песне с кодом 200
		} else {
			logger.Warnf("Song not found: group=%s, song=%s", group, song)
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Song not found"}) // Возвращаем ошибку 404, если песня не найдена
		}
	}
}
