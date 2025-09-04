package delete_newsletter

import (
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

		newsletterIDStr := chi.URLParam(r, "newsletter_id")
		newsletterID, err := strconv.ParseInt(newsletterIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid newsletter ID", http.StatusBadRequest)
			return
		}

		err = h.adminModule.Actions.DeleteNewsteller.Do(ctx, newsletterID)
		if err != nil {
			http.Error(w, "failed to delete newsletter: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
