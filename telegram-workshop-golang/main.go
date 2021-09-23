// Объявление пакета main
package main

// Подключение библиотек и пакетов
import (
	"fmt" // Библиотека для работы с потоками ввода/вывода и форматирования строк
	"log" // Библиотека для записи в лог приложения

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api" // API Telegram (добавлен синоним tgbotapi)
)

func main() {
	// API ключ для Telegram бота
	bot, err := tgbotapi.NewBotAPI("<KEY>")
	// Обработка ошибки при подключении к боту
	if err != nil {
		// Panic - вызовет ошибку, которая будет считаться недопустимой для дальнейшей работы программы
		log.Panic(fmt.Sprintf("Authorized error! Check your API key and connection!\n %s", err))
	} else {
		// Если подключение удалось - говорим пользователю, что мы авторизовались
		log.Printf("Authorized on account %s", bot.Self.FirstName)
	}

	/*
		Создание reply-клавиатуры
		Необходимые клавишы:
		1. О банке🏢 - рассказ о МТС Банке
		2. Личный кабинет😀 - отправка ссылки на личный кабинет МТС Банка (http://cb.mtsbank.ru/)
		3. Салон📍 - отправка геолокации салона МТС Банка в Кургане (широта = 55.442224, долгота = 65.346463)
		4. Курс валют📈 - отправка курса валют из Центробанка (http://www.cbr.ru/scripts/XML_daily.asp)

	*/

	// TODO: создать reply-клавиатуру

	// Добавление обновления чата
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	// Обработка пришедщих обновлений чата (записываются в переменную updates)
	for update := range updates {
		/*
			Создание обработчика обработчика событий
			Команды для обработки:
			1. /start - приветсвие при подписке к боту
			2. /stop - отписка от бота
			3-6. Клавиши из reply-клавиатуры
		*/

		// TODO: создать обработку пришедших команд
	}
}
