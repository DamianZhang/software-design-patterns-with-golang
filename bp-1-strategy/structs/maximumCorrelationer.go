package structs

type MaximumCorrelationer struct{}

func NewMaximumCorrelationer() *MaximumCorrelationer {
	return &MaximumCorrelationer{}
}

func (m *MaximumCorrelationer) Correlation(attributeBaseMatchees []*Individual) (fitMatchee *Individual) {
	return attributeBaseMatchees[len(attributeBaseMatchees)-1]
}
