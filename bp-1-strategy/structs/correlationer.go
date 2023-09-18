package structs

type Correlationer interface {
	Correlation(attributeBaseMatchees []*Individual) (fitMatchee *Individual)
}
