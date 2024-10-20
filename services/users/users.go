package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Yom3n/RecipeApiGo/db/database"
	"github.com/Yom3n/RecipeApiGo/utils"
	"github.com/google/uuid"
)

type Handler struct {
	db *database.Queries
}

func NewHandler(db *database.Queries) Handler {
	return Handler{db: db}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /users/", h.createUserHandler)
}

func (h *Handler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	jsonDecoder := json.NewDecoder(r.Body)
	params := parameters{}
	jsonDecoder.Decode(&params)
	if params.Name == "" {
		utils.RespondWithError(w, 400, "name is required")
		return
	}

	user, err := h.db.CreateUser(r.Context(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      params.Name,
		})
	if err != nil {
		utils.RespondWithError(w, 500, err.Error())
		return
	}
	utils.RespondWithJson(w, 201, user)
}
