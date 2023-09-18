package basic

type RelationshipAnalyzer interface {
	Parse(script string)
	GetMutualFriends(name1, name2 string) []string
}
