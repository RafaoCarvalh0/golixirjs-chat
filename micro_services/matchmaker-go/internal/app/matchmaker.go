package matchmaker

import (
	"matchmaker-go/internal/domain"

	"github.com/gin-gonic/gin"
)

func MatchUser(context *gin.Context, userWaitingQueue domain.UserQueue, queuedUserIds domain.QueuedUserIDs) {
	user, err := getUserFromContext(context)
	if err != nil {
		context.JSON(404, gin.H{"status": "user not found"})
		return
	}

	_, status, err := user.CreateMatch(userWaitingQueue, queuedUserIds)
	if err != nil {
		context.JSON(500, gin.H{"status": err.Error()})
	}

	context.JSON(200, gin.H{"status": status})
	return
}

func getUserFromContext(_ *gin.Context) (domain.User, error) {
	// TODO: fetch user from jwt
	return domain.User{ID: "foo"}, nil
}
