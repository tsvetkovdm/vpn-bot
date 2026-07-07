package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tsvetkovdm/vpn-bot/internal/config"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("не удалось загрузить конфиг: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatalf("не удалось создать клиента бота: %v", err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.Command() == "start" {
			msg.Text = "Привет, это твой личный VPN, для свободного сёрфинга в интернете"
		} else {
			msg.Text = "Неизвестная команда"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Println("не удалось отправить сообщение пользователю:", err)
		}
	}
}
