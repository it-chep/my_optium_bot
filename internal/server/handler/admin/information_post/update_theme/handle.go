package update_theme

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"net/http"
	"strconv"
)

type Handler struct {
	adminModule *admin.Module
}

func NewHandler(adminModule *admin.Module) *Handler {
	return &Handler{
		adminModule: adminModule,
	}
}

func (h *Handler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		themeIDStr := chi.URLParam(r, "theme_id")
		themeID, err := strconv.ParseInt(themeIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid post ID", http.StatusBadRequest)
			return
		}

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode request: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.adminModule.Actions.UpdateTheme.Do(ctx, themeID, req.Name, req.IsRequired)
		if err != nil {
			http.Error(w, "failed to get user data: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
