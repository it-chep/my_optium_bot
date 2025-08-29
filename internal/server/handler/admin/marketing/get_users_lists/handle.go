package get_users_lists

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

		lists, err := h.adminModule.Actions.GetUsersLists.Do(ctx)
		if err != nil {
			http.Error(w, "failed to get users lists: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := h.prepareResponse(lists)

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(lists []dto.UsersList) Response {
	return Response{
		Lists: lo.Map(lists, func(item dto.UsersList, _ int) List {
			return List{
				ID:         item.ID,
				Name:       item.Name,
				UsersCount: item.UsersCount,
			}
		}),
	}
}
