package app

import (
	"matchmaker-go/internal/core"
	"matchmaker-go/internal/domain"
)

type MatchmakerService struct {
	matchmaker *core.Matchmaker
}

func NewMatchmakerService(queue domain.UserQueue) *MatchmakerService {
	return &MatchmakerService{
		matchmaker: core.NewMatchmaker(queue),
	}
}

func (s *MatchmakerService) Match(user domain.User) (domain.Match, string, error) {
	return s.matchmaker.Match(user)
}
