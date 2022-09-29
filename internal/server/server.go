package server

import (
	"groupie-tracker/pkg/delivery"
	"net/http"
)

func NewServer(handler delivery.Handler) *http.Server {
	mux := http.NewServeMux()
	handler.InitRoutes(mux)
	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
