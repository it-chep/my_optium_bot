package tg_bot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewTgBot() (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	// todo: пока хз
	hook, err := tgbotapi.NewWebhook("https://четочето/telegram-webhook")
	if err != nil {
		return nil, err
	}

	_, err = bot.Request(hook)
	if err != nil {
		return nil, err
	}

	_, err = bot.GetWebhookInfo()
	if err != nil {
		return nil, err
	}

	return &Bot{bot: bot}, nil
}
