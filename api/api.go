package api

import (
	"fmt"
	"net/http"

	healthz "github.com/Yom3n/RecipeApiGo/services"
)

type APIServer struct {
	address string
	handler *http.ServeMux
}

func NewAPIServer(address string) APIServer {
	router := http.NewServeMux()
	healthzHandler := healthz.NewHandler()
	healthzHandler.RegisterRoutes(router)
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
