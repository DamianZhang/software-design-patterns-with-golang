package structs

import "fmt"

const (
	SymbolOfCharacterUp    string = "↑"
	SymbolOfCharacterDown  string = "↓"
	SymbolOfCharacterLeft  string = "←"
	SymbolOfCharacterRight string = "→"
	SymbolOfMonster        string = "M"
	SymbolOfTreasure       string = "x"
	SymbolOfObstacle       string = "□"
)

type AdventureGame struct {
	gameMap *GameMap
}

func NewAdventureGame() *AdventureGame {
	return &AdventureGame{
		gameMap: NewGameMap(),
	}
}

func (a *AdventureGame) Start() {
	var (
		stringsOfCharacter = []string{SymbolOfCharacterUp, SymbolOfCharacterDown, SymbolOfCharacterLeft, SymbolOfCharacterRight}
		stringsOfMonster   = []string{SymbolOfMonster}
	)

	fmt.Printf("adventure game is starting...\n\n")

	for a.gameMap.NumOfRolesByStrings(stringsOfCharacter) != 0 && a.gameMap.NumOfRolesByStrings(stringsOfMonster) != 0 {
		a.gameMap.PrintGameMap()
		a.gameMap.PrintCharacter()
		a.gameMap.LetRolesTakeTurn()
	}

	if a.gameMap.NumOfRolesByStrings(stringsOfMonster) == 0 {
		fmt.Println("〖主角獲勝〗")
	} else if a.gameMap.NumOfRolesByStrings(stringsOfCharacter) == 0 {
		fmt.Println("〖主角死亡〗")
	}
}
