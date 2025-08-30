package get_admin_messages

import (
	"encoding/json"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/samber/lo"
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
		messages, err := h.adminModule.Actions.GetAdminMessages.Do(ctx)
		if err != nil {
			http.Error(w, "failed to get user data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := h.prepareResponse(messages)

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(messages []dto.AdminMessage) Response {
	return Response{
		Messages: lo.Map(messages, func(item dto.AdminMessage, _ int) Message {
			return Message{
				ID:           item.ID,
				ScenarioName: item.ScenarioName,
				Type:         int8(item.Type),
				TypeName:     item.Type.String(),
				Text:         item.Message,
			}
		}),
	}
}
