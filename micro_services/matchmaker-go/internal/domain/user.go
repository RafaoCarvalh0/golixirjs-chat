package domain

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type User struct {
	ID string
}

type UserQueue map[string]User

type QueuedUserIDs []string

func (user *User) CreateMatch(queue UserQueue) (Match, string, error) {
	var status string
	var match Match

	user.addUserToQueue(queue)

	// TODO: add logic to consider user online status
	if len(queue) <= 1 {
		status = "waiting for a pair..."
		return match, status, nil
	}

	match, err := user.createRandomMatch(queue)
	if err != nil {
		return match, "match created", err
	}

	removeUsersFromQueue(queue, &match.User, &match.UserPair)

	return match, status, fmt.Errorf("unexpected error while trying to match users")
}

func (user *User) alreadyQueued(queue UserQueue) bool {
	if _, found := queue[user.ID]; found {
		return true
	}

	return false
}

func (user *User) addUserToQueue(queue UserQueue) {
	queue[user.ID] = *user
}

func (user *User) createRandomMatch(queue UserQueue) (Match, error) {
	var match Match

	randomUser, err := user.getRandomUserFromQueue(queue)
	if err != nil {
		return match, err
	}

	return Match{User: *user, UserPair: randomUser}, nil
}

func removeUsersFromQueue(queue UserQueue, users ...*User) {
	for _, user := range users {
		delete(queue, user.ID)
	}
}

func (user *User) getRandomUserFromQueue(queue UserQueue) (User, error) {
	var randomUser User
	queuedUsersCount := len(queue)

	if queuedUsersCount <= 1 {
		return randomUser, fmt.Errorf("no pair available")
	}

	var queuedUserIds []string
	for key := range queue {
		queuedUserIds = append(queuedUserIds, key)
	}

	randomUser = user.randomUserPair(queue, queuedUsersCount, queuedUserIds)

	return randomUser, nil
}

func (user *User) randomUserPair(queue UserQueue, queuedUsersCount int, queuedUserIds []string) User {
	var randomUser User

	randomUserId := queuedUserIds[randomIndex(queuedUsersCount)]
	if randomUser = queue[randomUserId]; randomUser.ID == user.ID {
		return user.randomUserPair(queue, queuedUsersCount, queuedUserIds)
	}

	return randomUser
}

func randomIndex(listLength int) int {
	randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(listLength)))
	return int(randomIndex.Int64())
}
