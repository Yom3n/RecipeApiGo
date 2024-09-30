package api

import (
	"fmt"
	"net/http"
)

type APIServer struct {
	address string
}

func NewAPIServer(address string) APIServer {
	return APIServer{
		address: address,
	}
}

func (s *APIServer) Run() error {
	fmt.Println("Starting server at :8080")
	server := &http.Server{
		Addr: s.address,
	}
	return server.ListenAndServe()
}
