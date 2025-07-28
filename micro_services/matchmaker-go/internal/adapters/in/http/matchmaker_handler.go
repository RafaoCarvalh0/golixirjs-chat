package http

import (
	"fmt"
	"matchmaker-go/internal/app"
	"matchmaker-go/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MatchmakerHandler struct {
	service *app.MatchmakerService
}

func NewMatchmakerHandler(queue domain.UserQueue) *MatchmakerHandler {
	return &MatchmakerHandler{
		service: app.NewMatchmakerService(queue),
	}
}

func (h *MatchmakerHandler) HandleMatchmaking(c *gin.Context) {
	user, err := getUserFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"data": gin.H{"match": nil, "message": err.Error(), "error": true}})
	}

	match, status, err := h.service.Match(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"data": gin.H{"match": nil, "message": err.Error(), "error": true}})
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"match": match, "message": status, "error": false}})
}

func getUserFromContext(context *gin.Context) (domain.User, error) {
	userID, exists := context.Get("userID")
	if !exists {
		return domain.NewUser(""), fmt.Errorf("user not found")
	}

	return domain.NewUser(userID.(string)), nil
}
