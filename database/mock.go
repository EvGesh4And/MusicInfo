package database

import "MusicInfo/models"

// mockDatabase - пример базы данных песен.
// Это структура данных, где ключом является название группы,
// а значением является карта, в которой ключом является название песни
// и значением — детали песни (ReleaseDate, Text, Link).
var mockDatabase = map[string]map[string]models.SongDetail{
	"Queen": {
		"Bohemian Rhapsody": {
			ReleaseDate: "16.07.2006",
			Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
			Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		},
		"Don't Stop Me Now": {
			ReleaseDate: "29.01.1979",
			Text:        "Tonight I'm gonna have myself a real good time,\nI feel alive.\nAnd the world, I'll turn it inside out, yeah.\nI'm floating around in ecstasy.",
			Link:        "https://www.youtube.com/watch?v=HgzGwKwLmgM",
		},
	},
	"The Beatles": {
		"Hey Jude": {
			ReleaseDate: "26.08.1968",
			Text:        "Hey Jude, don't make it bad.\nTake a sad song and make it better.\nRemember to let her into your heart,\nThen you can start to make it better.",
			Link:        "https://www.youtube.com/watch?v=8D6g1gR5mDo",
		},
		"Let It Be": {
			ReleaseDate: "06.03.1970",
			Text:        "When I find myself in times of trouble,\nMother Mary comes to me.\nSpeaking words of wisdom, let it be.",
			Link:        "https://www.youtube.com/watch?v=4xK5o4mMHiI",
		},
	},
	"Pink Floyd": {
		"Wish You Were Here": {
			ReleaseDate: "15.09.1975",
			Text:        "So, so you think you can tell\nHeaven from Hell?\nBlue skies from pain?\nCan you tell a green field\nFrom a cold steel rail?",
			Link:        "https://www.youtube.com/watch?v=IXdNnw99_4A",
		},
		"Comfortably Numb": {
			ReleaseDate: "30.11.1979",
			Text:        "Hello? Is there anybody in there?\nJust nod if you can hear me.\nIs there anyone home?",
			Link:        "https://www.youtube.com/watch?v=3j8w9t5U39k",
		},
	},
}

// GetSongDetail - получение деталей песни из "базы данных".
// Эта функция принимает название группы и название песни в качестве параметров.
// Если песня найдена, возвращает детали песни и true, иначе возвращает пустую структуру и false.
//
// @Param group query string true "Название группы"
// @Param song query string true "Название песни"
// @Success 200 {object} models.SongDetail "Детали песни"
// @Failure 404 {object} models.SongDetail "Ошибка: песня не найдена"
func GetSongDetail(group, song string) (models.SongDetail, bool) {
	if details, exists := mockDatabase[group][song]; exists {
		return details, true // Песня найдена
	}
	return models.SongDetail{}, false // Песня не найдена
}
