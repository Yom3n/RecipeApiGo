package api

import (
	"fmt"
	"net/http"
)

type APIServer struct {
	address string
	handler *http.ServeMux
}

func NewAPIServer(address string) APIServer {
handler := http.NewServeMux()
	handler.HandleFunc("GET /healthz/", healthz.HandlerReadines)
	return APIServer{
		address: address,
		handler: handler,
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
