package get_scenario_steps

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/samber/lo"
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
			http.Error(w, "invalid list ID", http.StatusBadRequest)
			return
		}

		steps, err := h.adminModule.Actions.GetScenarioSteps.Do(ctx, scenarioID)
		if err != nil {
			http.Error(w, "failed to update list: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := h.prepareResponse(steps)

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(steps []dto.Step) Response {
	return Response{
		Steps: lo.Map(steps, func(item dto.Step, _ int) Step {
			return Step{
				ID:        item.ID,
				Name:      item.Content,
				StepOrder: item.StepOrder,
			}
		}),
	}
}
