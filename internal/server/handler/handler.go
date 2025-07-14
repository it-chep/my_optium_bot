package handler

import (
	"fmt"
	"net/http"
)

type Handler struct {
	// todo: вероятно чето типо интерфейса будет с одним двумя методами (принимаем урл или контент тайп либо строку че тип ответил)
	// bot svc
	// admin svc
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) HandleAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(fmt.Sprintf("произошло восстановление системы: %v", err))
			}
		}()

		switch r.URL.Path {
		case "/telegram-webhook":
			h.bot().ServeHTTP(w, r)
		case "/admin":
			h.admin().ServeHTTP(w, r)
		default:
			_, _ = w.Write([]byte("По-моему, ты перепутал"))
		}
	}
}

func (h Handler) bot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Test"))
	}
}

func (h Handler) admin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Test"))
	}
}
