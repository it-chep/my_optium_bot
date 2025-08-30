package delete_admin_message

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
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

		messageIDStr := chi.URLParam(r, "message_id")
		messageID, err := strconv.ParseInt(messageIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid message ID", http.StatusBadRequest)
			return
		}

		var req Request
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode request: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.adminModule.Actions.DeleteAdminMessage.Do(ctx, messageID, dto.AdminType(req.Type))
		if err != nil {
			http.Error(w, "failed to delete admin message: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
