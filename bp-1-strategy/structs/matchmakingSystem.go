package structs

type MatchmakingSystem struct {
	matchType MatchType
}

func NewMatchmakingSystem(matchType MatchType) *MatchmakingSystem {
	return &MatchmakingSystem{matchType: matchType}
}

func (m *MatchmakingSystem) Matchmaking(matchmaker *Individual, matchees []*Individual) (fitMatchee *Individual) {
	return m.matchType.Matchmaking(matchmaker, matchees)
}
