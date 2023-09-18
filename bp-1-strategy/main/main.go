package main

import (
	"bp-1-strategy/structs"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("matchmaking system is starting...\n\n")

	var (
		matchmaker = NewMatchmaker()
		matchees   = NewRandomMatchees()
	)
	fmt.Printf("matchmaker: %v\n\n", matchmaker)

	var (
		distanceBase         = structs.NewDistanceBase()
		minimumCorrelationer = structs.NewMinimumCorrelationer()

		// habitsBase           = structs.NewHabitsBase()
		// maximumCorrelationer = structs.NewMaximumCorrelationer()
	)

	var (
		distanceBaseMinimumCorrelationer = structs.NewIndividualAttributeBaseStrategy(distanceBase, minimumCorrelationer)
		matchmakingSystem                = structs.NewMatchmakingSystem(distanceBaseMinimumCorrelationer)

		// distanceBaseMaximumCorrelationer = structs.NewIndividualAttributeBaseStrategy(distanceBase, maximumCorrelationer)
		// matchmakingSystem                = structs.NewMatchmakingSystem(distanceBaseMaximumCorrelationer)

		// habitsBaseMinimumCorrelationer = structs.NewIndividualAttributeBaseStrategy(habitsBase, minimumCorrelationer)
		// matchmakingSystem              = structs.NewMatchmakingSystem(habitsBaseMinimumCorrelationer)

		// habitsBaseMaximumCorrelationer = structs.NewIndividualAttributeBaseStrategy(habitsBase, maximumCorrelationer)
		// matchmakingSystem              = structs.NewMatchmakingSystem(habitsBaseMaximumCorrelationer)
	)

	fitMatchee := matchmakingSystem.Matchmaking(matchmaker, matchees)
	fmt.Printf("\nfitMatchee: %v\n", fitMatchee)
}

func NewMatchmaker() (matchmaker *structs.Individual) {
	var (
		habits          = []string{"打籃球", "煮菜", "玩遊戲"}
		coord           = structs.NewCoord(1.0, 1.0)
		individualProps = structs.NewIndividualProps("MALE", 18, "nice man", habits, coord)
	)
	matchmaker = structs.NewIndividual(individualProps)
	return matchmaker
}

func NewRandomMatchees() (matchees []*structs.Individual) {
	var (
		habitOptions      = []string{"吃美食", "打棒球", "旅遊", "聽音樂", "踢足球", "爬山", "衝浪", "打籃球", "煮菜", "玩遊戲"}
		lenOfHabitOptions = len(habitOptions)
	)

	for i := 2; i <= 10; i++ {
		habits := make([]string, 0)
		for j := 1; j <= 3; j++ {
			habits = append(habits, habitOptions[rand.Intn(lenOfHabitOptions)])
		}

		var (
			coord           = structs.NewCoord(float64(i), float64(i))
			individualProps = structs.NewIndividualProps("MALE", 18, "nice man", habits, coord)
			matchee         = structs.NewIndividual(individualProps)
		)
		matchees = append(matchees, matchee)
	}

	rand.Shuffle(len(matchees), func(i int, j int) {
		matchees[i], matchees[j] = matchees[j], matchees[i]
	})

	return matchees
}
