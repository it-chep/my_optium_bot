package get_users

import (
	"encoding/json"
	"fmt"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/samber/lo"
	"net/http"
	"time"
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

		usersDTO, err := h.adminModule.Actions.GetUsers.Do(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := h.prepareResponse(usersDTO)
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(usersDTO []dto.User) Response {
	return Response{
		Users: lo.Map(usersDTO, func(usr dto.User, _ int) User {
			return User{
				ID:          usr.ID,
				Name:        usr.FullName,
				TgID:        usr.TgID,
				Sex:         usr.Sex.String(),
				MetricsLink: usr.MetricsLink,
				Birthday:    usr.BirthDate.Format(time.DateOnly),
			}
		}),
	}
}
