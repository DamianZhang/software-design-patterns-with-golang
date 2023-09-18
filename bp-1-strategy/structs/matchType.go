package structs

type MatchType interface {
	Matchmaking(matchmaker *Individual, matchees []*Individual) (fitMatchee *Individual)
}
