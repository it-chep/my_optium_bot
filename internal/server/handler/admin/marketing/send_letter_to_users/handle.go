package send_letter_to_users

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

		newsletterIDStr := chi.URLParam(r, "newsletters_id")
		newsletterID, err := strconv.ParseInt(newsletterIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid list ID", http.StatusBadRequest)
			return
		}

		err = h.adminModule.Actions.SendLetterToUsers.Do(ctx, newsletterID)
		if err != nil {
			http.Error(w, "failed to get user data: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
