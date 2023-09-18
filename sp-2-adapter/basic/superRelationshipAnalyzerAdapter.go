package basic

import (
	"fmt"
	"strings"
)

type SuperRelationshipAnalyzerAdapter struct {
	superRelationshipAnalyzer *SuperRelationshipAnalyzer
}

func NewSuperRelationshipAnalyzerAdapter() *SuperRelationshipAnalyzerAdapter {
	return &SuperRelationshipAnalyzerAdapter{
		superRelationshipAnalyzer: NewSuperRelationshipAnalyzer(),
	}
}

func (s *SuperRelationshipAnalyzerAdapter) Parse(script string) {
	s.superRelationshipAnalyzer.Init(s.convertClientScriptToSuperRelationshipAnalyzerScript(script))
}

func (s *SuperRelationshipAnalyzerAdapter) convertClientScriptToSuperRelationshipAnalyzerScript(clientScript string) (superRelationshipAnalyzerScript string) {
	members := strings.Split(clientScript, "\n")

	for indexOfMember, member := range members {
		var (
			infoOfMember           = strings.Split(member, ": ")
			nameOfMember           = infoOfMember[0]
			namesOfFriendsOfMember = strings.Split(infoOfMember[1], " ")
		)

		for indexOfNameOfFriendOfMember, nameOfFriendOfMember := range namesOfFriendsOfMember {
			if indexOfMember == len(members)-1 && indexOfNameOfFriendOfMember == len(namesOfFriendsOfMember)-1 {
				superRelationshipAnalyzerScript += fmt.Sprintf("%s -- %s", nameOfMember, nameOfFriendOfMember)
			} else {
				superRelationshipAnalyzerScript += fmt.Sprintf("%s -- %s\n", nameOfMember, nameOfFriendOfMember)
			}
		}
	}

	return superRelationshipAnalyzerScript
}

func (s *SuperRelationshipAnalyzerAdapter) GetMutualFriends(name1, name2 string) []string {
	var (
		mutualFriends             = make([]string, 0)
		superRelationshipAnalyzer = s.superRelationshipAnalyzer
		members                   = superRelationshipAnalyzer.Members()
		friendsOfName1AndName2    = append(members[name1], members[name2]...)
	)

	for _, friendOfName1AndName2 := range friendsOfName1AndName2 {
		if superRelationshipAnalyzer.IsMutualFriend(friendOfName1AndName2, name1, name2) &&
			!sliceHasElement(mutualFriends, friendOfName1AndName2) {
			mutualFriends = append(mutualFriends, friendOfName1AndName2)
		}
	}

	return mutualFriends
}

func sliceHasElement(slice []string, element string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}

	return false
}
