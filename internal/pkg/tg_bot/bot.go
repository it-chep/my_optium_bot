package tg_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
)

type Config interface {
	WebhookURL() string
	Token() string
}

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewTgBot(cfg Config) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Token())
	if err != nil {
		return nil, err
	}

	hook, _ := tgbotapi.NewWebhook(cfg.WebhookURL() + cfg.Token() + "/")
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

func (b *Bot) HandleUpdate(r *http.Request) (*tgbotapi.Update, error) {
	return b.bot.HandleUpdate(r)
}
