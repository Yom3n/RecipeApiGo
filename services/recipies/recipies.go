package recipies

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Yom3n/RecipeApiGo/auth"
	"github.com/Yom3n/RecipeApiGo/db/db"
	"github.com/Yom3n/RecipeApiGo/utils"
	"github.com/google/uuid"
)

type RecipiesHandler struct {
	db *db.Queries
}

func NewHandler(db *db.Queries) RecipiesHandler {
	return RecipiesHandler{
		db: db,
	}
}

func (h *RecipiesHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST recipies/", h.HandleCreateRecipe)

}

func (h *RecipiesHandler) HandleCreateRecipe(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		utils.RespondWithError(w, 401, err.Error())
		return
	}
	jsonDecoder := json.NewDecoder(r.Body)
	type RecipeInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	payload := RecipeInput{}
	err = jsonDecoder.Decode(&payload)
	if err != nil {
		utils.RespondWithError(w, 500, err.Error())
		return
	}
	h.db.CreateRecipe(r.Context(), db.CreateRecipeParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       payload.Title,
		Description: payload.Description,
		AuthorID:    uuid.New(), // TODO Regenerate schema + sqlc and provide user id
	})
}
