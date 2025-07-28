package main

import (
	"matchmaker-go/internal/adapters/in/http"
	"matchmaker-go/internal/adapters/in/http/middleware"
	"matchmaker-go/internal/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	var waitingQueue domain.UserQueue

	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(middleware.JWTAuthMiddleware())

	handler := http.NewMatchmakerHandler(waitingQueue)
	router.POST("/create-match", handler.HandleMatchmaking)

	router.Run(":8000")
}
