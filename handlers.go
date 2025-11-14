package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// структура товара
type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Desc  string `json:"desc"`
}

var Products []Product

// загрузка JSON
func LoadProducts() {
	data, err := ioutil.ReadFile("data/products.json")
	if err != nil {
		log.Fatal("Ошибка чтения products.json:", err)
	}

	err = json.Unmarshal(data, &Products)
	if err != nil {
		log.Fatal("Ошибка JSON:", err)
	}

	fmt.Println("Products loaded:", len(Products))
}

func GetProductByID(id string) *Product {
	for _, p := range Products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

// основная обработка callback'ов
func HandleCallbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	data := update.CallbackQuery.Data

	switch {

	// главное меню
	case data == "menu":
		bot.Send(MenuCategories(chatID, msgID))

	// кнопка назад
	case data == "back_start":
		bot.Send(StartMenu(chatID))

	// страница товара
	case len(data) > 8 && data[:8] == "product_":
		id := data[8:]
		p := GetProductByID(id)
		bot.Send(ProductPage(chatID, msgID, p))

	// платёжная страница
	case len(data) > 4 && data[:4] == "buy_":
		id := data[4:]
		p := GetProductByID(id)
		bot.Send(PaymentPage(chatID, msgID, p))

	// пользователь нажал "I Paid"
	case len(data) > 5 && data[:5] == "paid_":
		id := data[5:]
		p := GetProductByID(id)

		bot.Send(tgbotapi.NewMessage(chatID,
			"Спасибо! Ваш платёж отправлен на проверку."))

		// уведомление админу
		user := update.CallbackQuery.From

		adminMsg := fmt.Sprintf(
			"🟢 Новый платёж!\n\n"+
				"Пользователь: @%s (ID: %d)\n"+
				"Товар: %s\n"+
				"Сумма: %d RUB",
			user.UserName, user.ID, p.Name, p.Price,
		)

		msg := tgbotapi.NewMessage(AdminID, adminMsg)
		bot.Send(msg)
	}
}
