package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	LoadProducts()

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		// /start
		if update.Message != nil && update.Message.Text == "/start" {
			bot.Send(StartMenu(update.Message.Chat.ID))
			continue
		}

		if update.CallbackQuery == nil {
			continue
		}

		chatID := update.CallbackQuery.Message.Chat.ID
		msgID := update.CallbackQuery.Message.MessageID
		data := update.CallbackQuery.Data

		switch {

		case data == "menu":
			bot.Send(MenuCategories(chatID, msgID))

		case data == "back_start":
			bot.Send(StartMenu(chatID))

		case len(data) > 8 && data[:8] == "product_":
			id := data[8:]
			p := GetProductByID(id)
			bot.Send(ProductPage(chatID, msgID, p))

		case len(data) > 4 && data[:4] == "buy_":
			id := data[4:]
			p := GetProductByID(id)
			bot.Send(PaymentPage(chatID, msgID, p))

		case len(data) > 5 && data[:5] == "paid_":
			id := data[5:]
			p := GetProductByID(id)

			bot.Send(tgbotapi.NewMessage(chatID, "Спасибо! Платёж отправлен на проверку."))

			// уведомление админу
			adminMsg := tgbotapi.NewMessage(AdminID,
				"🟢 Новый платёж!\n\n"+
					"Пользователь: @"+update.CallbackQuery.From.UserName+"\n"+
					"Товар: "+p.Name+"\n"+
					"Сумма: "+fmt.Sprint(p.Price)+" RUB",
			)
			bot.Send(adminMsg)
		}
	}
}
