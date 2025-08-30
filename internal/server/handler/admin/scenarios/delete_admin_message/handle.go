package delete_admin_message

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
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
		//ctx := r.Context()
		//
		//messageIDStr := chi.URLParam(r, "message_id")
		//messageID, err := strconv.ParseInt(messageIDStr, 10, 64)
		//if err != nil {
		//	http.Error(w, "invalid message ID", http.StatusBadRequest)
		//	return
		//}
		//
		//h.adminModule.Actions.DeleteAdminMessage.Do(ctx, messageID)
	}
}
