package recipies

import (
	"encoding/json"
	"log"
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
	router.HandleFunc("POST /recipies/", h.HandleCreateRecipe)
	router.HandleFunc("GET /recipies/", h.HandleGetUserRecipies)
}

func (h *RecipiesHandler) HandleCreateRecipe(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		utils.RespondWithError(w, 401, err.Error())
		log.Println(err.Error())
		return
	}
	user, err := h.db.GetUserByApiKey(r.Context(), string(apiKey))
	if err != nil {
		utils.RespondWithError(w, 500, err.Error())
		log.Println(err.Error())
		return
	}
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
	err = jsonDecoder.Decode(&payload)
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

func (h *RecipiesHandler) HandleGetUserRecipies(w http.ResponseWriter, r *http.Request) {
	apiKey, _ := auth.GetApiKey(r.Header)
	user, _ := h.db.GetUserByApiKey(r.Context(), string(apiKey))
	recipies, _ := h.db.GetUserRecipies(r.Context(), user.ID)
	utils.RespondWithJson(w, 200, recipies)
}
