package middleware

import (
	"log"
	"net/http"
)

// type authorizedRouteHandler func(w http.ResponseWriter, r *http.Request)

// func AuthorizedRoute(w http.ResponseWriter, r *http.Request) func(w http.ResponseWriter, r *http.Request, user db.User) {
// 	apiKey, err := auth.GetApiKey(r.Header)
// 	if err != nil || len(apiKey) == 0 {
// 		utils.RespondWithError(w, 403, "Unauthorized access")
// 		return nil
// 	}
// }

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
