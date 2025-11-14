package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	LoadConfig()
	LoadProducts()

	go func() {
		if err := runBot(); err != nil {
			log.Fatalf("bot stopped: %v", err)
		}
	}()

	startHTTPServer()
}

func runBot() error {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		return fmt.Errorf("bot init: %w", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.Text == "/start" {
			if _, err := bot.Send(StartMenu(update.Message.Chat.ID)); err != nil {
				log.Printf("send start menu: %v", err)
			}
			continue
		}

		if update.CallbackQuery != nil {
			HandleCallbacks(bot, update)
			continue
		}
	}

	return nil
}
