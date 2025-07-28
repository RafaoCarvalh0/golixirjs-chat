package main

import (
	"matchmaker-go/internal/adapters/in/http"
	"matchmaker-go/internal/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	var waitingQueue domain.UserQueue

	router := gin.Default()
	router.Use(gin.Recovery())

	handler := http.NewMatchmakerHandler(waitingQueue)
	router.POST("/matchmake", handler.HandleMatchmaking)

	router.Run(":8000")
}
