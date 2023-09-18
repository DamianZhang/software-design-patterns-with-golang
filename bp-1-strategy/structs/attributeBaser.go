package structs

type AttributeBaser interface {
	AttributeBase(matchmaker *Individual, matchees []*Individual) (attributeBaseMatchees []*Individual)
}
