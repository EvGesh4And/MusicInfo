# MusicInfo API

## Описание

MusicInfo предоставляет API для получения информации о песнях на основе названия группы и песни. Проект был создан для проверки корректности работы основной программы https://github.com/EvGesh4And/MusicLibrary.

## Версия

1.0

## Контактная информация

- **Имя**: Евгений
- **Email**: [i@evgesh4.ru](mailto:i@evgesh4.ru)

## Запуск приложения

Для запуска приложения выполните следующие шаги:

1. Клонируйте репозиторий:

    ```bash
    git clone https://github.com/EvGesh4And/MusicInfo
    cd MusicInfo
    ```

2. Установите необходимые зависимости:

    ```bash
    go mod tidy
    ```

3. Настройте переменные окружения. Создайте файл `.env` в корне проекта и добавьте следующие параметры:

    ```plaintext
    API_PORT=9090  # Опционально, для настройки порта API
    ```

4. Запустите приложение:

    ```bash
    go run main.go
    ```

Сервер будет запущен на порту, указанном в переменной окружения `API_PORT` (по умолчанию — 9090).

## Использование API

API MusicInfo поддерживает следующие эндпоинты:

### Получение информации о песне
- **URL**: `/info`
- **Метод**: `GET`
- **Параметры запроса**:
  - `group` (обязательный): название группы
  - `song` (обязательный): название песни
- **Ответ**:
  - `200 OK`: информация о песне
  - `400 Bad Request`: отсутствуют необходимые параметры запроса
  - `404 Not Found`: песня не найдена
  - `500 Internal Server Error`: внутренняя ошибка сервера

## База данных

Ниже представлен список групп и их песни, для которых можно получить дополнительную информацию:

#### Queen
1. **Bohemian Rhapsody**
2. **Don't Stop Me Now**
#### The Beatles
1. **Hey Jude**
2. **Let It Be**
#### Pink Floyd
1. **Wish You Were Here**
2. **Comfortably Numb**

## Логирование

Приложение логирует ошибки и запросы с использованием стандартного логирования Go.

## Swagger Documentation

Документация по API доступна по следующему адресу:

- [Swagger Documentation](http://localhost:9090/swagger/index.html)
