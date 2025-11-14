package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Модель товара
type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Desc  string `json:"desc"`
}

var Products []Product

// Загружаем JSON
func LoadProducts() {
	data, err := ioutil.ReadFile("data/products.json")
	if err != nil {
		log.Fatal("Ошибка чтения products.json:", err)
	}

	err = json.Unmarshal(data, &Products)
	if err != nil {
		log.Fatal("Ошибка JSON:", err)
	}
}

func GetProductByID(id string) *Product {
	for _, p := range Products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func HandleCallbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	data := update.CallbackQuery.Data

	switch {

	case data == "menu":
		bot.Send(MenuCategories(chatID))

	case data == "back_start":
		bot.Send(StartMenu(chatID))

	case len(data) > 8 && data[:8] == "product_":
		id := data[8:]
		p := GetProductByID(id)
		bot.Send(ProductPage(chatID, p))

	case len(data) > 4 && data[:4] == "buy_":
		id := data[4:]
		p := GetProductByID(id)
		bot.Send(PaymentPage(chatID, p))

	case len(data) > 5 && data[:5] == "paid_":
		id := data[5:]
		p := GetProductByID(id)

		// Пользователь
		bot.Send(tgbotapi.NewMessage(chatID, "Спасибо! Платёж отправлен на проверку."))

		// Админу
		msg := tgbotapi.NewMessage(AdminID,
			fmt.Sprintf("🟢 Новый платёж!\n\nПользователь: @%s\nТовар: %s\nСумма: %d RUB",
				update.CallbackQuery.From.UserName,
				p.Name,
				p.Price,
			),
		)
		bot.Send(msg)
	}
}
