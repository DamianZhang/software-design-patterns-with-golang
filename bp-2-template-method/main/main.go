package main

import (
	"bp-2-template-method/showdown"
	"bp-2-template-method/uno"
	"fmt"
)

func main() {
	gameA, err := showdown.NewGame(NewShowdownPlayers())
	if err != nil {
		fmt.Println(err)
	}
	gameA.Start()

	// gameB, err := uno.NewGame(NewUNOPlayers())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// gameB.Start()
}

func NewShowdownPlayers() []*showdown.Player {
	players := make([]*showdown.Player, 0)

	for i := 1; i <= 4; i++ {
		player := showdown.NewPlayer()
		players = append(players, player)
	}

	return players
}

func NewUNOPlayers() []*uno.Player {
	players := make([]*uno.Player, 0)

	for i := 1; i <= 4; i++ {
		player := uno.NewPlayer()
		players = append(players, player)
	}

	return players
}
