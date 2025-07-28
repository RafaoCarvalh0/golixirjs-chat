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

type RandomMatcher interface {
	randomMatch(queue UserQueue) (Match, error)
}

func (user *User) CreateMatch(queue UserQueue) (Match, string, error) {
	return CreateMatch(user, queue, user)
}

func CreateMatch(user *User, queue UserQueue, matcher RandomMatcher) (Match, string, error) {
	var match Match

	status := "waiting for a pair..."

	user.addUserToQueue(queue)

	// TODO: add logic to consider user online status
	if len(queue) <= 1 {
		return match, status, nil
	}

	match, err := matcher.randomMatch(queue)
	if err != nil {
		return match, status, fmt.Errorf("unexpected error while trying to match users")
	}

	removeUsersFromQueue(queue, &match.User, &match.UserPair)

	return match, "match created", err
}

func (user *User) addUserToQueue(queue UserQueue) {
	queue[user.ID] = *user
}

func (user *User) randomMatch(queue UserQueue) (Match, error) {
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
		if key != user.ID {
			queuedUserIds = append(queuedUserIds, key)
		}
	}

	randomUserId := queuedUserIds[randomIndexFromList(queuedUserIds)]

	return queue[randomUserId], nil
}

func randomIndexFromList[T any](list []T) int {
	listLength := len(list)

	randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(listLength)))
	return int(randomIndex.Int64())
}
