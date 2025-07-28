package core

import (
	"matchmaker-go/internal/domain"
)

type Matchmaker struct {
	Queue domain.UserQueue
}

func NewMatchmaker(queue domain.UserQueue) *Matchmaker {
	return &Matchmaker{Queue: queue}
}

func (m *Matchmaker) Match(user domain.User) (domain.Match, string, error) {
	return user.CreateMatch(m.Queue)
}
