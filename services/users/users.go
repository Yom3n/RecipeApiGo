package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Yom3n/RecipeApiGo/db/db"
	"github.com/Yom3n/RecipeApiGo/utils"
	"github.com/google/uuid"
)

type UsersHandler struct {
	db *db.Queries
}

func NewHandler(db *db.Queries) UsersHandler {
	return UsersHandler{db: db}
}

func (h *UsersHandler) RegisterRoutes(handler *http.ServeMux) {
	handler.HandleFunc("POST /users/", h.createUserHandler)
}

func (h *UsersHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
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
		db.CreateUserParams{
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
