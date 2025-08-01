package tg_bot

import (
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"github.com/samber/lo"
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

func (b *Bot) GetUser(message dto.Message) (bot_dto.User, error) {
	member, err := b.bot.GetChatMember(tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID: message.ChatID,
			UserID: message.User,
		},
	})
	if err != nil {
		return bot_dto.User{}, err
	}

	return bot_dto.User{
		ID:       member.User.ID,
		Name:     member.User.FirstName,
		UserName: member.User.UserName,
		IsAdmin:  member.IsCreator() || member.IsAdministrator(),
	}, nil
}

func (b *Bot) SendMessage(msg bot_dto.Message) error {
	message := tgbotapi.NewMessage(msg.Chat, msg.Text)

	if len(msg.Buttons) != 0 {
		message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				lo.Map(msg.Buttons, func(b dto.StepButton, _ int) tgbotapi.InlineKeyboardButton {
					return tgbotapi.NewInlineKeyboardButtonData(b.Text, b.Text)
				})...,
			),
		)
	}
	_, err := b.bot.Send(message)
	return err
}

func (b *Bot) SendMessages(messages []bot_dto.Message) error {
	for _, msg := range messages {
		if err := b.SendMessage(msg); err != nil {
			return err
		}
	}

	return nil
}
