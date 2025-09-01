package get_post_by_id

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

		postIDStr := chi.URLParam(r, "post_id")
		postID, err := strconv.ParseInt(postIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid post ID", http.StatusBadRequest)
			return
		}

		post, err := h.adminModule.Actions.GetPostByID.Do(ctx, postID)
		if err != nil {
			http.Error(w, "failed to get post data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := h.prepareResponse(post)

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(post dto.InformationPost) Response {
	return Response{
		Post: Post{
			ID:            post.ID,
			Name:          post.Name,
			ThemeName:     post.ThemeName,
			ThemeID:       post.PostsThemeID,
			Order:         post.OrderInTheme,
			Message:       post.PostText,
			MediaID:       post.MediaID,
			ContentTypeID: post.ContentTypeID,
		},
	}
}
