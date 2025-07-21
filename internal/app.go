package internal

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot"
	"log"
	"log/slog"

	"github.com/it-chep/my_optium_bot.git/internal/config"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/server"
	"github.com/jackc/pgx/v5"
)

type App struct {
	logger *slog.Logger
	config *config.Config
	conn   *pgx.Conn

	server *server.Server
	bot    *tg_bot.Bot

	botService *bot.Bot
}

func New(ctx context.Context) *App {
	cfg := config.NewConfig()

	app := &App{
		config: cfg,
	}

	app.initDB(ctx).
		initTgBot(ctx).
		initServer(ctx)

	return app
}

func (a *App) Run(context.Context) {
	a.logger.Info("start server")
	log.Fatal(a.server.ListenAndServe())
}
