package internal

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/config"
	"log/slog"
	"net/http"
)

type App struct {
	logger *slog.Logger
	config *config.Config
	server *http.Server
}

func New(ctx context.Context) *App {
	cfg := config.NewConfig()

	app := &App{
		config: cfg,
	}
	return app
}

func (app *App) Run(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			app.logger.Error("application recovered from panic", slog.Any("error", r))
		}
	}()

	app.logger.Info("start server")
	err := app.server.ListenAndServe()
	if err != nil {
		app.logger.Error("Не удалось запустить приложение")
		return
	}
}
