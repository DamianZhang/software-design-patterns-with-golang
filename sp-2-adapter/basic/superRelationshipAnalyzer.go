package basic

import (
	"strings"
)

type SuperRelationshipAnalyzer struct {
	members map[string][]string // key = nameOfMember, value = namesOfFriendsOfMember
}

func NewSuperRelationshipAnalyzer() *SuperRelationshipAnalyzer {
	return &SuperRelationshipAnalyzer{
		members: make(map[string][]string),
	}
}

func (s *SuperRelationshipAnalyzer) Init(script string) {
	relationships := strings.Split(script, "\n")

	for _, relationship := range relationships {
		var (
			names = strings.Split(relationship, " -- ")
			name1 = names[0]
			name2 = names[1]
		)

		s.members[name1] = append(s.members[name1], name2)
	}
}

func (s *SuperRelationshipAnalyzer) IsMutualFriend(targetName, name2, name3 string) bool {
	return s.hasTargetFriend(targetName, name2) && s.hasTargetFriend(targetName, name3)
}

func (s *SuperRelationshipAnalyzer) hasTargetFriend(targetName, name string) bool {
	namesOfFriendsOfMember, IsExisting := s.members[name]
	if !IsExisting {
		return false
	}

	for _, nameOfFriendOfMember := range namesOfFriendsOfMember {
		if nameOfFriendOfMember == targetName {
			return true
		}
	}

	return false
}

func (s *SuperRelationshipAnalyzer) Members() map[string][]string {
	return s.members
}
