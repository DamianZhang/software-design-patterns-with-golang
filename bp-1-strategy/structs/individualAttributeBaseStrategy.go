package structs

type IndividualAttributeBaseStrategy struct {
	attributeBaser AttributeBaser
	correlationer  Correlationer
}

func NewIndividualAttributeBaseStrategy(attributeBaser AttributeBaser, correlationer Correlationer) *IndividualAttributeBaseStrategy {
	return &IndividualAttributeBaseStrategy{
		attributeBaser: attributeBaser,
		correlationer:  correlationer,
	}
}

func (i *IndividualAttributeBaseStrategy) Matchmaking(matchmaker *Individual, matchees []*Individual) (fitMatchee *Individual) {
	attributeBaseMatchees := i.attributeBaser.AttributeBase(matchmaker, matchees)
	return i.correlationer.Correlation(attributeBaseMatchees)
}
