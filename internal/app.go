package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/it-chep/my_optium_bot.git/internal/config"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/server"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	config *config.Config
	pool   *pgxpool.Pool

	server *server.Server
	bot    *tg_bot.Bot

	modules Modules
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
		initServer(ctx)

	return app
}

func (a *App) Run(context.Context) {
	fmt.Println("start server")
	log.Fatal(a.server.ListenAndServe())
}
