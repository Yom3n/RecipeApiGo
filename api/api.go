package api

import (
	"fmt"
	"net/http"

	"github.com/Yom3n/RecipeApiGo/db/db"
	"github.com/Yom3n/RecipeApiGo/middleware"
	healthz "github.com/Yom3n/RecipeApiGo/services/healthz"
	"github.com/Yom3n/RecipeApiGo/services/recipies"
	"github.com/Yom3n/RecipeApiGo/services/users"
)

type APIServer struct {
	address string
	handler http.Handler
}

func NewAPIServer(address string, db *db.Queries) APIServer {
	handler := http.NewServeMux()

	healthzHandler := healthz.NewHandler()
	healthzHandler.RegisterRoutes(handler)

	usersHandler := users.NewHandler(db)
	usersHandler.RegisterRoutes(handler)

	recipiesHandler := recipies.NewHandler(db)
	recipiesHandler.RegisterRoutes(handler)

	return APIServer{
		address: address,
		handler: handler,
	}
}

func (s *APIServer) Run() error {
	fmt.Println("Starting server at :8080")
	server := &http.Server{
		Addr:    s.address,
		Handler: middleware.NewLoggerMiddleware(s.handler),
	}
	return server.ListenAndServe()
}
