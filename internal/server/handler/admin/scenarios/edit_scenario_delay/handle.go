package edit_scenario_delay

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

		scenarioIDStr := chi.URLParam(r, "scenario_id")
		scenarioID, err := strconv.ParseInt(scenarioIDStr, 10, 64)
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

		err = h.adminModule.Actions.EditScenarioDelay.Do(ctx, scenarioID, req.Delay)
		if err != nil {
			http.Error(w, "failed to update scenario delay: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
