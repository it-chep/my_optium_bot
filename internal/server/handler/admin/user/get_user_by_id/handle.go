package get_user_by_id

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	actionDto "github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/get_user_by_id/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/samber/lo"
	"net/http"
	"strconv"
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

		userIDStr := chi.URLParam(r, "user_id")
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid user ID", http.StatusBadRequest)
			return
		}

		baseData, err := h.adminModule.Actions.GetUserByID.Do(ctx, userID)
		if err != nil {
			http.Error(w, "failed to get user data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := h.prepareResponse(baseData)

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(data actionDto.Response) Response {
	return Response{
		User: UserData{
			ID:          data.UserData.ID,
			Name:        data.UserData.FullName,
			Sex:         data.UserData.Sex.String(),
			TgID:        data.UserData.TgID,
			MetricsLink: data.UserData.MetricsLink,
			Birthday:    data.UserData.BirthDate.Format(time.DateOnly),
		},
		Lists: lo.Map(data.Lists, func(item dto.UsersList, _ int) List {
			return List{
				ID:   item.ID,
				Name: item.Name,
			}
		}),
		Posts: lo.Map(data.Posts, func(item dto.InformationPost, _ int) Post {
			return Post{
				ID:              item.ID,
				Name:            item.Name,
				IsRequiredTheme: item.ThemeIsRequired,
			}
		}),
		Scenarios: lo.Map(data.Scenarios, func(item dto.PatientScenario, _ int) Scenario {
			return Scenario{
				ID:        item.ScenarioID,
				Name:      dto.ScenarioNameMap[item.ScenarioID],
				NextDelay: item.ScheduledTime.String(),
			}
		}),
	}
}
