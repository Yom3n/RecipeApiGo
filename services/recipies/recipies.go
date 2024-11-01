package recipies

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Yom3n/RecipeApiGo/db/db"
	"github.com/Yom3n/RecipeApiGo/middleware"
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

func (h *RecipiesHandler) RegisterRoutes(handler *http.ServeMux) {
	handler.Handle("POST /recipies/", middleware.NewEnsureAuth(h.HandleCreateRecipe, h.db))
	handler.Handle("GET /recipies/", middleware.NewEnsureAuth(h.HandleGetUserRecipies, h.db))
}

func (h *RecipiesHandler) HandleCreateRecipe(w http.ResponseWriter, r *http.Request, user db.User) {
	if r.Body == http.NoBody {
		utils.RespondWithError(w, 400, "Body can't be empty")
		return
	}
	jsonDecoder := json.NewDecoder(r.Body)
	type RecipeInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	payload := RecipeInput{}
	err := jsonDecoder.Decode(&payload)
	if err != nil {
		utils.RespondWithError(w, 500, err.Error())
		log.Println(err.Error())
		return
	}
	if len(payload.Title) == 0 {
		utils.RespondWithError(w, 400, "title is required")
		return
	}
	if len(payload.Description) == 0 {
		utils.RespondWithError(w, 400, "desciprion is required")
		return
	}
	recipe, err := h.db.CreateRecipe(r.Context(), db.CreateRecipeParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       payload.Title,
		Description: payload.Description,
		AuthorID:    user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 500, err.Error())
		return
	}
	utils.RespondWithJson(w, 201, recipe)
}

func (h *RecipiesHandler) HandleGetUserRecipies(w http.ResponseWriter, r *http.Request, user db.User) {
	recipies, _ := h.db.GetUserRecipies(r.Context(), user.ID)
	utils.RespondWithJson(w, 200, recipies)
}
