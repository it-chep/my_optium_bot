package update_list_name

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

		listIDStr := chi.URLParam(r, "list_id")
		listID, err := strconv.ParseInt(listIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid list ID", http.StatusBadRequest)
			return
		}

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode request: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.adminModule.Actions.UpdateListName.Do(ctx, listID, req.Name)
		if err != nil {
			http.Error(w, "failed to update list: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
