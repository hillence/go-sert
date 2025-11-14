package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartMenu(chatID int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, "🔍 Welcome to bot!")

	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Личный сертификат", "menu"),
		),
	)

	msg.ReplyMarkup = kb
	return msg
}

func MenuCategories(chatID int64, msgID int) tgbotapi.EditMessageTextAndMarkup {
	text := "Select a tariff or category from the list below 👇"

	kb := tgbotapi.NewInlineKeyboardMarkup()

	// Добавляем товары из JSON
	for _, p := range Products {
		kb.InlineKeyboard = append(kb.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(p.Name, "product_"+p.ID),
			),
		)
	}

	// Кнопка назад
	kb.InlineKeyboard = append(kb.InlineKeyboard,
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Back", "back_start"),
		),
	)

	return tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, text, kb)
}

func ProductPage(chatID int64, msgID int, p *Product) tgbotapi.EditMessageTextAndMarkup {
	text := fmt.Sprintf(
		"<b>%s</b>\nPrice: %d RUB\n\n%s",
		p.Name, p.Price, p.Desc,
	)

	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Приобрести", "buy_"+p.ID),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Back", "menu"),
		),
	)

	msg := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, text, kb)
	msg.ParseMode = "HTML"
	return msg
}

func PaymentPage(chatID int64, msgID int, p *Product) tgbotapi.EditMessageTextAndMarkup {

	text := fmt.Sprintf(`
<b>%s БАНК</b>

Payment Method: Приобрести
Amount to pay: <b>%d RUB</b>

Переведите по номеру карты:
<b>%s</b>

%s  
%s
`,
		CardBank, p.Price, CardNumber, CardBank, CardName)

	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅ I Paid.", "paid_"+p.ID),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("❌ Cancel Payment", "menu"),
		),
	)

	msg := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, text, kb)
	msg.ParseMode = "HTML"
	return msg
}
