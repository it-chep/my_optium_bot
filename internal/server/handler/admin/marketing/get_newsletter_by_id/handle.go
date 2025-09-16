package get_newsletter_by_id

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

		newsletterIDStr := chi.URLParam(r, "newsletters_id")
		newsletterID, err := strconv.ParseInt(newsletterIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid newsletter ID", http.StatusBadRequest)
			return
		}

		newsletters, err := h.adminModule.Actions.GetNewsletterByID.Do(ctx, newsletterID)
		if err != nil {
			http.Error(w, "failed to get newsletter data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := h.prepareResponse(newsletters)

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(newsletter dto.Newsletter) Response {
	return Response{
		NewsLetter{
			ID:   newsletter.ID,
			Name: newsletter.Name,

			StatusID:   int8(newsletter.StatusID),
			StatusName: newsletter.StatusID.String(),

			UsersCount: newsletter.RecipientsCount,

			Text:        newsletter.Text,
			UsersLists:  newsletter.UsersLists,
			MediaID:     newsletter.MediaID,
			ContentType: int8(newsletter.ContentType),
		},
	}
}
