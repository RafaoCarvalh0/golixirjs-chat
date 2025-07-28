package router

import (
	"matchmaker-go/internal/adapters/in/http"
	"matchmaker-go/internal/adapters/in/http/middleware"
	"matchmaker-go/internal/domain"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	waitingQueue := make(domain.UserQueue)

	handler := http.NewMatchmakerHandler(waitingQueue)

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.JWTAuthMiddleware())

	router.POST("/create-match", handler.HandleMatchmaking)

	return router
}
