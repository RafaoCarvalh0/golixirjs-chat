package main

import (
	"log"
	"matchmaker-go/internal/router"
)

func main() {
	router := router.NewRouter()

	if err := router.Run(":8000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
