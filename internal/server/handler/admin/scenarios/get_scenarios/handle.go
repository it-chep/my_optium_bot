package get_scenarios

import (
	"encoding/json"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/samber/lo"
	"net/http"
	"strings"
	"unicode/utf8"
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

		scenarios, err := h.adminModule.Actions.GetScenarios.Do(ctx)
		if err != nil {
			http.Error(w, "failed to get scenarios data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := h.prepareResponse(scenarios)

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(scenarios []dto.Scenario) Response {
	scens := lo.Filter(scenarios, func(item dto.Scenario, _ int) bool {
		if strings.Contains(item.Description, "Не используется") ||
			utf8.RuneCountInString(item.Description) == 0 {
			return false
		}
		return true
	})
	return Response{
		Scenarios: lo.Map(scens, func(item dto.Scenario, _ int) Scenario {
			return Scenario{
				ID:    item.ID,
				Name:  item.Name,
				Delay: item.Delay.String(),
			}
		}),
	}
}
