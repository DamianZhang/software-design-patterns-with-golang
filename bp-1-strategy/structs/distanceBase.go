package structs

import (
	"fmt"
	"math"
	"sort"
)

type DistanceBase struct{}

func NewDistanceBase() *DistanceBase {
	return &DistanceBase{}
}

func calculateDistance(CoordOfMatchmaker, CoordOfMatchee *Coord) int {
	differenceOfX := math.Pow(CoordOfMatchmaker.X()-CoordOfMatchee.X(), 2)
	differenceOfY := math.Pow(CoordOfMatchmaker.Y()-CoordOfMatchee.Y(), 2)
	return int(math.Pow(differenceOfX+differenceOfY, 0.5))
}

func (d *DistanceBase) AttributeBase(matchmaker *Individual, matchees []*Individual) (attributeBaseMatchees []*Individual) {
	attributeBaseMatchees = matchees

	sort.SliceStable(attributeBaseMatchees, func(i, j int) bool {
		distanceOfI := calculateDistance(matchmaker.Coord(), attributeBaseMatchees[i].Coord())
		distanceOfJ := calculateDistance(matchmaker.Coord(), attributeBaseMatchees[j].Coord())
		if distanceOfI != distanceOfJ {
			return distanceOfI < distanceOfJ
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
