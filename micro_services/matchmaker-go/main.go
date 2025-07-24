package main

import (
	matchmaker "matchmaker-go/internal/app"
	"matchmaker-go/internal/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	var waitingQueue domain.UserQueue
	var queuedUserIds domain.QueuedUserIDs

	router := gin.Default()

	router.Use(gin.Recovery())

	router.POST("/matchmake", func(context *gin.Context) {
		matchmaker.MatchUser(context, waitingQueue, queuedUserIds)
	})

	router.Run(":8000")
}
