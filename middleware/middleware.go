package middleware

import (
	"log"
	"net/http"

	"github.com/Yom3n/RecipeApiGo/auth"
	"github.com/Yom3n/RecipeApiGo/db/db"
	"github.com/Yom3n/RecipeApiGo/utils"
)

type AuthenticationHandler func(http.ResponseWriter, *http.Request, db.User)

type EnsureAuthenticated struct {
	Handler AuthenticationHandler
	db      *db.Queries
}

func (ea *EnsureAuthenticated) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil || len(apiKey) == 0 {
		utils.RespondWithError(w, 403, "Unauthorized access")
		return
	}
	user, err := ea.db.GetUserByApiKey(r.Context(), string(apiKey))
	if err != nil {
		log.Println(err.Error())
		utils.RespondWithError(w, 500, "Internal error")
	}
	ea.Handler(w, r, user)
}

func NewEnsureAuth(handlerToWrap AuthenticationHandler, db *db.Queries) *EnsureAuthenticated {
	return &EnsureAuthenticated{Handler: handlerToWrap, db: db}
}

type LoggerMiddleware struct {
	Handler http.Handler
}

func NewLoggerMiddleware(handlerToWrap http.Handler) *LoggerMiddleware {
	return &LoggerMiddleware{Handler: handlerToWrap}
}

func (lm *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	url := r.URL
	log.Printf("%s on %s", method, url)
	lm.Handler.ServeHTTP(w, r)
}
