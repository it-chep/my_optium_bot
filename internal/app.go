package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/it-chep/my_optium_bot.git/internal/config"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/worker"
	"github.com/it-chep/my_optium_bot.git/internal/server"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/lo"
)

type Workers []worker.Worker

type App struct {
	config *config.Config
	pool   *pgxpool.Pool

	server *server.Server
	bot    *tg_bot.Bot

	modules Modules
	workers Workers
}

type Modules struct {
	Bot *bot.Bot
}

func New(ctx context.Context) *App {
	cfg := config.NewConfig()

	app := &App{
		config: cfg,
	}

	app.initDB(ctx).
		initTgBot(ctx).
		initModules(ctx).
		initJobs(ctx).
		initServer(ctx)

	return app
}

func (a *App) Run(ctx context.Context) {
	fmt.Println("start server")
	ctx = logger.ContextWithLogger(ctx, logger.New())
	for _, w := range a.workers {
		w.Start(ctx)
	}
	if !a.config.UseWebhook() {
		fmt.Println("Режим поллинга")
		// Режим поллинга
		go func() {
			for update := range a.bot.GetUpdates() {
				if update.ChatMember != nil {
					if lo.Contains([]string{"left", "kicked"}, update.ChatMember.NewChatMember.Status) {
						return
					}
					usrID := update.ChatMember.NewChatMember.User.ID
					chat := update.ChatMember.Chat.ID
					_ = a.modules.Bot.Actions.InvitePatient.InvitePatient(ctx, usrID, chat)
				}

				if update.FromChat() == nil || update.SentFrom() == nil {
					return
				}
				logger.Message(ctx, "Обработка ивента")

				txt := ""
				mediaID := ""
				if update.Message != nil {
					txt = update.Message.Text
					// фото
					if update.Message.Photo != nil {
						// массив фото разбивает фотографию на 4 качества, берем самое плохое )
						mediaID = update.Message.Photo[0].FileID
					}
					// видео
					if update.Message.Video != nil {
						mediaID = update.Message.Video.FileID
					}
					// документ
					if update.Message.Document != nil {
						mediaID = update.Message.Document.FileID
					}
					// кружок
					if update.Message.VideoNote != nil {
						mediaID = update.Message.VideoNote.FileID
					}
					// голосовое сообщение
					if update.Message.Voice != nil {
						mediaID = update.Message.Voice.FileID
					}
					// аудио сообщение
					if update.Message.Audio != nil {
						mediaID = update.Message.Audio.FileID
					}
				} else if update.CallbackQuery != nil {
					txt = update.CallbackQuery.Data
				}

				msg := dto.Message{
					User:    update.SentFrom().ID,
					Text:    txt,
					ChatID:  update.FromChat().ID,
					MediaID: mediaID,
				}
				err := a.modules.Bot.Route(ctx, msg)

				if err != nil {
					logger.Error(ctx, "Ошибка при обработке ивента", err)
				}
			}
		}()
	}

	log.Fatal(a.server.ListenAndServe())
}
