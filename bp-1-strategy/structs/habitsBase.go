package structs

import (
	"fmt"
	"sort"
)

type HabitsBase struct{}

func NewHabitsBase() *HabitsBase {
	return &HabitsBase{}
}

func numOfIntersections(habitsOfMatchmaker, habitsOfMatchee []string) (numOfIntersections int) {
	for _, habitOfMatchmaker := range habitsOfMatchmaker {
		for _, habitOfMatchee := range habitsOfMatchee {
			if habitOfMatchmaker == habitOfMatchee {
				numOfIntersections++
			}
		}
	}
	return numOfIntersections
}

func (h *HabitsBase) AttributeBase(matchmaker *Individual, matchees []*Individual) (attributeBaseMatchees []*Individual) {
	attributeBaseMatchees = matchees

	sort.SliceStable(attributeBaseMatchees, func(i, j int) bool {
		numOfIntersectionsOfI := numOfIntersections(matchmaker.Habits(), attributeBaseMatchees[i].Habits())
		numOfIntersectionsOfJ := numOfIntersections(matchmaker.Habits(), attributeBaseMatchees[j].Habits())
		if numOfIntersectionsOfI != numOfIntersectionsOfJ {
			return numOfIntersectionsOfI < numOfIntersectionsOfJ
		}

		idOfI := attributeBaseMatchees[i].Id()
		idOfJ := attributeBaseMatchees[j].Id()
		if idOfI != idOfJ {
			return idOfI < idOfJ
		}

		return false
	})

	for i, matchee := range attributeBaseMatchees {
		fmt.Printf("matchee %d: %v\n", i, matchee)
	}

	return attributeBaseMatchees
}
