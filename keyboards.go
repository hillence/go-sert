package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Главное меню
func StartMenu(chatID int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, "🔍 Добро пожаловать!")

	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Личный сертификат", "menu"),
		),
	)

	msg.ReplyMarkup = kb
	return msg
}

// Категории
func MenuCategories(chatID int64) tgbotapi.MessageConfig {
	text := "Выберите товар 👇"

	kb := tgbotapi.NewInlineKeyboardMarkup()

	for _, p := range Products {
		kb.InlineKeyboard = append(kb.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(p.Name, "product_"+p.ID),
			),
		)
	}

	kb.InlineKeyboard = append(kb.InlineKeyboard,
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", "back_start"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	return msg
}

// Страница товара
func ProductPage(chatID int64, p *Product) tgbotapi.MessageConfig {
	text := fmt.Sprintf("<b>%s</b>\nЦена: %d RUB\n\n%s", p.Name, p.Price, p.Desc)

	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Приобрести", "buy_"+p.ID),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", "menu"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = kb
	msg.ParseMode = "HTML"
	return msg
}

// Страница оплаты
func PaymentPage(chatID int64, p *Product) tgbotapi.MessageConfig {
	text := fmt.Sprintf(`
<b>%s БАНК</b>

Сумма к оплате: <b>%d RUB</b>

Переведите по номеру карты:
<b>%s</b>

%s
%s
`, CardBank, p.Price, CardNumber, CardBank, CardName)

	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅ Я оплатил", "paid_"+p.ID),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("❌ Отменить", "menu"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = kb
	msg.ParseMode = "HTML"
	return msg
}
