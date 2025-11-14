package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	LoadProducts()

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Fatal("Bot start error:", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		// Команда /start
		if update.Message != nil && update.Message.Text == "/start" {
			bot.Send(StartMenu(update.Message.Chat.ID))
			continue
		}

		// Обработка кнопок
		if update.CallbackQuery != nil {
			HandleCallbacks(bot, update)
			continue
		}
	}
}
