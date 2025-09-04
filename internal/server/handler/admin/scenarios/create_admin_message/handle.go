package create_admin_message

import (
	"encoding/json"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	actionDto "github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/create_admin_message/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"net/http"
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

		if r.Method != http.MethodPost {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode request: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err := h.adminModule.Actions.CreateAdminMessage.Do(ctx, actionDto.CreateMessageRequest{
			Message:    req.Message,
			Type:       dto.AdminType(req.Type),
			ScenarioID: req.ScenarioID,
			StepOrder:  req.StepOrder,
		})

		if err != nil {
			http.Error(w, "failed to create admin message: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
