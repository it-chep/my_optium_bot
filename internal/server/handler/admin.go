package handler

import "net/http"

func (h *Handler) admin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Test"))
	}
}
