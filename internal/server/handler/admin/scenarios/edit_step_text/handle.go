package edit_step_text

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

		stepIDStr := chi.URLParam(r, "step_id")
		stepID, err := strconv.ParseInt(stepIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid step ID", http.StatusBadRequest)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}

		var req Request
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode request: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.adminModule.Actions.EditStepText.Do(ctx, stepID, req.Message)
		if err != nil {
			http.Error(w, "failed to update step message: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
