package update_newsletter

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/update_newsletter/dto"
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

		newsletterIDStr := chi.URLParam(r, "newsletters_id")
		newsletterID, err := strconv.ParseInt(newsletterIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid newsletter ID", http.StatusBadRequest)
			return
		}

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode request: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.adminModule.Actions.UpdateNewsletter.Do(ctx, newsletterID, dto.UpdateNewsletterDTO{
			Name:          req.Name,
			Text:          req.Text,
			UsersLists:    req.UsersLists,
			MediaID:       req.MediaID,
			ContentTypeID: req.ContentTypeID,
		})
		if err != nil {
			http.Error(w, "failed to update newsletter: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
