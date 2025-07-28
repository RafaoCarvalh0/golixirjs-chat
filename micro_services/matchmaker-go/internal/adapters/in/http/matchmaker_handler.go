package http

import (
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
		c.JSON(http.StatusNotFound, gin.H{"data": gin.H{"match": nil, "message": "user not found", "error": true}})
		return
	}

	match, status, err := h.service.Match(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": gin.H{"match": nil, "message": err.Error(), "error": true}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"match": match, "message": status, "error": false}})
}

func getUserFromContext(_ *gin.Context) (domain.User, error) {
	// TODO: fetch user from JWT
	return domain.User{ID: "foo"}, nil
}
