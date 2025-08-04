package domain

import "sync"

func NewUserQueue() UserQueue {
	return UserQueue{
		mu:      &sync.Mutex{},
		userMap: make(map[string]User),
	}
}

func NewUser(id string) User {
	return User{
		ID: id,
	}
}
