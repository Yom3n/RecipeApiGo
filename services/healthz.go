package healthz

import (
	"net/http"

	"github.com/Yom3n/RecipeApiGo/utils"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /healthz/", h.HandlerReadines)
}

func (h *Handler) HandlerReadines(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, struct{}{})
}
