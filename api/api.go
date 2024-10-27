package api

import (
	"fmt"
	"net/http"

	"github.com/Yom3n/RecipeApiGo/db/db"
	healthz "github.com/Yom3n/RecipeApiGo/services/healthz"
	"github.com/Yom3n/RecipeApiGo/services/recipies"
	"github.com/Yom3n/RecipeApiGo/services/users"
)

type APIServer struct {
	address string
	handler *http.ServeMux
}

func NewAPIServer(address string, db *db.Queries) APIServer {
	router := http.NewServeMux()

	healthzHandler := healthz.NewHandler()
	healthzHandler.RegisterRoutes(router)

	usersHandler := users.NewHandler(db)
	usersHandler.RegisterRoutes(router)

	recipiesHandler := recipies.NewHandler(db)
	recipiesHandler.RegisterRoutes(router)

	return APIServer{
		address: address,
		handler: router,
	}
}

func (s *APIServer) Run() error {
	fmt.Println("Starting server at :8080")
	server := &http.Server{
		Addr:    s.address,
		Handler: s.handler,
	}
	return server.ListenAndServe()
}
