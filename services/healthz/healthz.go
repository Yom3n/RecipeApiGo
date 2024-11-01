package healthz

import (
	"net/http"

	"github.com/Yom3n/RecipeApiGo/utils"
)

type HealthzHandler struct{}

func NewHandler() *HealthzHandler {
	return &HealthzHandler{}
}

func (h *HealthzHandler) RegisterRoutes(handler *http.ServeMux) {
	handler.HandleFunc("GET /healthz/", h.HandlerReadines)
}

func (h *HealthzHandler) HandlerReadines(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, struct{}{})
}
