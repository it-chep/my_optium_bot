package create_newsletter

import (
	"encoding/json"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/create_newsletter/dto"
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

		err := h.adminModule.Actions.CreateNewsLetter.Do(ctx, dto.Request{
			Name:          req.Name,
			UsersList:     req.UsersList,
			Text:          req.Text,
			MediaID:       req.MediaID,
			ContentTypeID: req.ContentTypeID,
		})
		if err != nil {
			http.Error(w, "failed to create newsletter: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
