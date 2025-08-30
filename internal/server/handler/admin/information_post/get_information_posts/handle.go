package get_information_posts

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

		posts, err := h.adminModule.Actions.GetInformationPosts.Do(ctx)
		if err != nil {
			http.Error(w, "failed to get user data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := h.prepareResponse(posts)

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(posts []dto.InformationPostListView) Response {
	return Response{
		Posts: lo.Map(posts, func(item dto.InformationPostListView, _ int) Post {
			return Post{
				ID:              item.ID,
				Name:            item.Name,
				ThemeName:       item.PostThemeName,
				Order:           item.OrderInTheme,
				IsThemeRequired: item.ThemeIsRequired,
			}
		}),
	}
}
