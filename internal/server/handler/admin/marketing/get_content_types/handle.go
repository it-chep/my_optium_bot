package get_content_types

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

		contentTypes, err := h.adminModule.Actions.GetContentTypes.Do(ctx)
		if err != nil {
			http.Error(w, "failed to get contentTypes data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := h.prepareResponse(contentTypes)

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse(contentTypes []dto.ContentTypeDTO) Response {
	return Response{
		ContentTypes: lo.Map(contentTypes, func(item dto.ContentTypeDTO, _ int) ContentType {
			return ContentType{
				ID:   item.ID,
				Name: item.Name,
			}
		}),
	}
}
