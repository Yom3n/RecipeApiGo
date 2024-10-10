package healthz

import (
	"net/http"

	"github.com/Yom3n/RecipeApiGo/utils"
)

func HandlerReadines(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, nil)

}
