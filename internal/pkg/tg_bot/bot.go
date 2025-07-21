package tg_bot

import (
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config interface {
	WebhookURL() string
	Token() string
	UseWebhook() bool
}

type Bot struct {
	bot        *tgbotapi.BotAPI
	updates    tgbotapi.UpdatesChannel
	useWebhook bool
}

func NewTgBot(cfg Config) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Token())
	if err != nil {
		return nil, err
	}

	// Режим вебхуков
	if cfg.UseWebhook() {
		hook, _ := tgbotapi.NewWebhook(cfg.WebhookURL() + cfg.Token() + "/")
		_, err = bot.Request(hook)
		if err != nil {
			return nil, err
		}

		_, err = bot.GetWebhookInfo()
		if err != nil {
			return nil, err
		}

		return &Bot{
			bot:        bot,
			useWebhook: true,
		}, nil
	}

	// Режим поллинга
	_, _ = bot.Request(tgbotapi.DeleteWebhookConfig{})

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	return &Bot{
		bot:        bot,
		updates:    updates,
		useWebhook: false,
	}, nil
}

func (b *Bot) HandleUpdate(r *http.Request) (*tgbotapi.Update, error) {
	return b.bot.HandleUpdate(r)
}

func (b *Bot) GetUpdates() tgbotapi.UpdatesChannel {
	return b.updates
}
