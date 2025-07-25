package main

import (
	matchmaker "matchmaker-go/internal/app"
	"matchmaker-go/internal/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	var waitingQueue domain.UserQueue

	router := gin.Default()

	router.Use(gin.Recovery())

	router.POST("/matchmake", func(context *gin.Context) {
		matchmaker.NewMatch(context, waitingQueue)
	})

	router.Run(":8000")
}
