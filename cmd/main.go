package main

import (
	"groupie-tracker/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Printf("%v", err)
	}
}
