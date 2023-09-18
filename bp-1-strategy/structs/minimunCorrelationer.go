package structs

type MinumumCorrelationer struct{}

func NewMinimumCorrelationer() *MinumumCorrelationer {
	return &MinumumCorrelationer{}
}

func (m *MinumumCorrelationer) Correlation(attributeBaseMatchees []*Individual) (fitMatchee *Individual) {
	return attributeBaseMatchees[0]
}
