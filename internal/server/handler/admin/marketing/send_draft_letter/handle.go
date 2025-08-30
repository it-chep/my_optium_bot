package send_draft_letter

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
		//newsletterIDStr := chi.URLParam(r, "newsletters_id") // todo user_id
		//newsletterID, err := strconv.ParseInt(newsletterIDStr, 10, 64)
		//if err != nil {
		//	http.Error(w, "invalid list ID", http.StatusBadRequest)
		//	return
		//}
		//
		//h.adminModule.Actions.SentDraftLetter.Do(ctx, newsletterID)
	}
}
