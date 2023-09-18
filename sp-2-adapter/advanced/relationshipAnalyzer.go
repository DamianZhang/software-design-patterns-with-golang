package advanced

type RelationshipAnalyzer interface {
	Parse(script string) RelationshipGraph
}
