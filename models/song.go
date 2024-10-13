package models

// SongDetail - структура для представления деталей песни.
// @Description Детали песни, включая дату выхода, текст и ссылку
type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// ErrorResponse - структура для представления сообщения об ошибке.
// @Description Структура, используемая для возврата сообщений об ошибках.
type ErrorResponse struct {
	Error string `json:"error"` // Сообщение об ошибке
}
