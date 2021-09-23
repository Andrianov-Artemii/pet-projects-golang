// Объявление пакета main
package main

// Подключение библиотек и пакетов
import (
	"encoding/xml" // Библиотека для парсинга xml
	"fmt"          // Библиотека для работы с потоками ввода/вывода и форматирования строк
	"net/http"     // Библиотека для работы с http запросами

	"golang.org/x/net/html/charset" // Библиотека для подключения кодировок

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api" // API Telegram (добавлен синоним tgbotapi)
)

/* TODO: Hello-функция для бота
Входные параметры:
	1. Ссылка на бота из main.go (*tgbotapi.BotAPI)
	2. Копия объекта обновления чата из main.go (tgbotapi.Update)
	3 (опционально). Копия на reply-клавиатуру из main.go (tgbotapi.ReplyKeyboardMarkup)

Для создания сообщения исполуется метод tgbotapi.NewMessage(ID чата, строка-сообщение)
*/

/* TODO: Функция с описание бота (текст + картинкой)
Входные параметры:
	1. Ссылка на бота из main.go (*tgbotapi.BotAPI)
	2. Копия объекта обновления чата из main.go (tgbotapi.Update)

Для создания сообщения исполуется метод tgbotapi.NewPhotoUpload(ID чата, путь до файла*)
Для добавление подписи использовать переменную Caption (тип-текст)
* Файл лежит по пути src/about.jpg
*/

/* TODO: Функция, отправляющая ссылку на Личный кабинет пользоальтеля
Входные параметры:
	1. Ссылка на бота из main.go (*tgbotapi.BotAPI)
	2. Копия объекта обновления чата из main.go (tgbotapi.Update)

Для создания сообщения исполуется метод tgbotapi.NewMessage(ID чата, строка-сообщение)

Ссылка на Личный кабинет: http://cb.mtsbank.ru/
*/

/* TODO: Функция, отправляющая геолокацию отделения МТС Банка
Входные параметры:
	1. Ссылка на бота из main.go (*tgbotapi.BotAPI)
	2. Копия объекта обновления чата из main.go (tgbotapi.Update)

Для создания сообщения-подписи исполуется метод tgbotapi.NewMessage(ID чата, строка-сообщение)
Для создания сообщения с геолокации исполуется метод tgbotapi.NewLocation(ID чата, широта (число), долгота (число))

Вот одно из отделений:
	Широта 55.442224
	Долгота 65.346463
*/

/* TODO: Парсинг курса валют из Центробанка и отправка пользователю
Входные параметры:
	1. Ссылка на бота из main.go (*tgbotapi.BotAPI)
	2. Копия объекта обновления чата из main.go (tgbotapi.Update)

Для создания сообщения исполуется метод tgbotapi.NewMessage(ID чата, строка-сообщение)

Для http-запроса используется метод http.Get("<url>") - возвращает ответ от сервера (тело ответа лежит в поле Body) и ошибку

Для парсинга xml используется Decoder
decoder := xml.NewDecoder(<Тело для парсинга)
Для установки кодировки используется CharsetReader и пакет charset
decoder.CharsetReader = charset.NewReaderLabel
Вызов парсинга
decoder.Decode(<ссылка на объект>)
*/

/* TODO: Функция для неверной командой
Входные параметры:
	1. Ссылка на бота из main.go (*tgbotapi.BotAPI)
	2. Копия объекта обновления чата из main.go (tgbotapi.Update)

Для создания сообщения исполуется метод tgbotapi.NewMessage(ID чата, строка-сообщение)
*/

/* TODO: Bue-функция для бота
Входные параметры:
	1. Ссылка на бота из main.go (*tgbotapi.BotAPI)
	2. Копия объекта обновления чата из main.go (tgbotapi.Update)

Для создания сообщения исполуется метод tgbotapi.NewMessage(ID чата, строка-сообщение)
*/
