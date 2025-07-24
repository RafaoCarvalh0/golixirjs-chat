package matchmaker

import (
	"matchmaker-go/internal/domain"

	"github.com/gin-gonic/gin"
)

func NewMatch(context *gin.Context, userWaitingQueue domain.UserQueue, queuedUserIds domain.QueuedUserIDs) {
	user, err := getUserFromContext(context)
	if err != nil {
		context.JSON(404, gin.H{"data": gin.H{"match": nil, "message": "user not found", "error": true}})
		return
	}

	match, status, err := user.CreateMatch(userWaitingQueue, &queuedUserIds)
	if err != nil {
		context.JSON(500, gin.H{"data": gin.H{"match": nil, "message": err.Error(), "error": true}})
	}

	context.JSON(200, gin.H{"data": gin.H{"match": match, "message": status, "error": false}})
}

func getUserFromContext(_ *gin.Context) (domain.User, error) {
	// TODO: fetch user from jwt
	return domain.User{ID: "foo"}, nil
}
