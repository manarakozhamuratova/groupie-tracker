package app

import (
	"groupie-tracker/internal/server"
	d "groupie-tracker/pkg/delivery"
	"log"
)

func Run() error {
	handler := d.NewHandler()

	s := server.NewServer(*handler)
	log.Printf("http://localhost:8080/")
	return s.ListenAndServe()
}
