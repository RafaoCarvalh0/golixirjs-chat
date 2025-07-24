package domain

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type User struct {
	ID              string
	PositionInQueue *int
}

type UserQueue map[string]User

type QueuedUserIDs []string

func (user *User) CreateMatch(queue UserQueue, queuedUserIds *QueuedUserIDs) (Match, string, error) {
	var status string
	var match Match

	if !user.alreadyQueued(queue) {
		user.addUserToQueues(queue, queuedUserIds)
	}

	// TODO: add logic to consider user online status
	if len(*queuedUserIds) <= 1 {
		status = "waiting for a pair..."
		return match, status, nil
	}

	match, err := user.createRandomMatch(queue, queuedUserIds)
	if err != nil {
		return match, "match created", err
	}

	return match, status, fmt.Errorf("unexpected error while trying to match users")
}

func (user *User) alreadyQueued(queue UserQueue) bool {
	if _, found := queue[user.ID]; found {
		return true
	}

	return false
}

func (user *User) addUserToQueues(queue UserQueue, queuedUserIds *QueuedUserIDs) {
	userPositionInQueue := len(*queuedUserIds)

	queue[user.ID] = *user
	*queuedUserIds = append(*queuedUserIds, user.ID)

	user.PositionInQueue = new(int)
	*user.PositionInQueue = userPositionInQueue
}

func (user *User) createRandomMatch(queue UserQueue, queuedUserIds *QueuedUserIDs) (Match, error) {
	var match Match

	randomUser, err := user.getRandomUserFromQueue(queue, queuedUserIds)
	if err != nil {
		return match, err
	}

	return Match{User: *user, UserPair: randomUser}, nil
}

func (user *User) getRandomUserFromQueue(queue UserQueue, queuedUserIds *QueuedUserIDs) (User, error) {
	var randomUser User
	queuedUsersCount := len(*queuedUserIds)

	if queuedUsersCount <= 1 {
		return randomUser, fmt.Errorf("no pair available")
	}

	randomUserIdIndex, err := rand.Int(rand.Reader, big.NewInt(int64(queuedUsersCount-1)))
	if err != nil {
		return randomUser, fmt.Errorf("could not pick a random pair")
	}

	index := int(randomUserIdIndex.Int64())
	randomUserId := (*queuedUserIds)[index]
	if randomUser = queue[randomUserId]; randomUser.ID == user.ID {
		return user.getRandomUserFromQueue(queue, queuedUserIds)
	}

	delete(queue, user.ID)
	delete(queue, randomUser.ID)

	*queuedUserIds = append((*queuedUserIds)[:index], (*queuedUserIds)[index+1:]...)

	return randomUser, nil
}
