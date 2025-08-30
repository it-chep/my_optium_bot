package get_recepients_count

import (
	"encoding/json"
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

		recepientsCount, err := h.adminModule.Actions.GetRecepientsCount.Do(ctx, req.ListIDs)
		if err != nil {
			http.Error(w, "failed to create newsletter: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := h.prepareResponse(recepientsCount)
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(recepientsCount int64) Response {
	return Response{
		Count: recepientsCount,
	}
}
