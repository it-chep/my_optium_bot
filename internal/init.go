package internal

import (
	"context"
	"log"

	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/server"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler"
	"github.com/jackc/pgx/v5"
)

func init() {
	// ignore db columns that doesn't exist at the destination
	dbscanAPI, err := pgxscan.NewDBScanAPI(dbscan.WithAllowUnknownColumns(true))
	if err != nil {
		panic(err)
	}

	api, err := pgxscan.NewAPI(dbscanAPI)
	if err != nil {
		panic(err)
	}

	pgxscan.DefaultAPI = api
}

func (a *App) initDB(ctx context.Context) *App {
	conn, err := pgx.Connect(ctx, a.config.PgConn())
	if err != nil {
		log.Fatalf("[FATAL] не удалось создать кластер базы данных: %s", err)
	}

	a.conn = conn
	return a
}

func (a *App) initTgBot(context.Context) *App {
	bot, err := tg_bot.NewTgBot()
	if err != nil {
		log.Fatal(err)
	}
	a.bot = bot
	return a
}

func (a *App) initServer(context.Context) *App {
	// todo: в NewHandler передаем сервис для админки или бота
	h := handler.NewHandler()
	srv := server.New(h)
	a.server = srv
	return a
}
