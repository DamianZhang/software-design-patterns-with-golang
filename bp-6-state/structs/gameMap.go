package structs

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	RowOfGameMap           int = 3
	ColOfGameMap           int = 3
	InitialNumOfCharacters int = 1
	InitialNumOfMonsters   int = 2
	InitialNumOfTreasures  int = 5
	InitialNumOfObstacles  int = 1
)

var (
	Rander = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type GameMap struct {
	roles      []*Role
	mapObjects [][]IMapObject
}

func NewGameMap() *GameMap {
	g := &GameMap{
		roles:      make([]*Role, 0),
		mapObjects: makeMapObjects(),
	}

	g.initializeMapObjects(g.generateRandomCharacter, InitialNumOfCharacters)
	g.initializeMapObjects(g.generateRandomMonster, InitialNumOfMonsters)
	g.initializeMapObjects(g.generateRandomTreasure, InitialNumOfTreasures)
	g.initializeMapObjects(g.generateRandomObstacle, InitialNumOfObstacles)

	return g
}

func makeMapObjects() [][]IMapObject {
	mapObjects := make([][]IMapObject, RowOfGameMap)

	for rowIndex := range mapObjects {
		mapObjects[rowIndex] = make([]IMapObject, ColOfGameMap)
	}

	return mapObjects
}

func (g *GameMap) Roles() []*Role {
	return g.roles
}

func (g *GameMap) RemoveRole(role *Role) {
	for i, r := range g.roles {
		if r != nil && r.Coordinate() == role.Coordinate() {
			g.roles[i] = nil
			return
		}
	}
}

func (g *GameMap) GetMapObject(coordinate *Coordinate) IMapObject {
	row := coordinate.Row()
	col := coordinate.Col()

	return g.mapObjects[row][col]
}

func (g *GameMap) RemoveMapObject(coordinate *Coordinate) {
	row := coordinate.Row()
	col := coordinate.Col()

	g.mapObjects[row][col] = nil
}

func (g *GameMap) MoveMapObject(mapObject IMapObject, coordinate *Coordinate) {
	row := mapObject.Coordinate().Row()
	col := mapObject.Coordinate().Col()
	g.mapObjects[row][col] = nil
	mapObjectInformationBeforeMoving := fmt.Sprintf("%s(%d,%d)", mapObject.String(), col, row)

	mapObject.SetCoordinate(coordinate)

	row = coordinate.Row()
	col = coordinate.Col()
	g.mapObjects[row][col] = mapObject
	mapObjectInformationAfterMoving := fmt.Sprintf("%s(%d,%d)", mapObject.String(), col, row)

	fmt.Printf("%s 移動到 %s\n", mapObjectInformationBeforeMoving, mapObjectInformationAfterMoving)
}

func (g *GameMap) NumOfRolesByStrings(stringsOfRole []string) (numOfRoles int) {
	for _, role := range g.roles {
		if role != nil {
			for _, stringOfRole := range stringsOfRole {
				switch role.String() {
				case stringOfRole:
					numOfRoles++
				}
			}
		}
	}

	return numOfRoles
}

func (g *GameMap) PrintCharacter() {
	for _, role := range g.roles {
		if role != nil {
			switch role.String() {
			case SymbolOfCharacterUp, SymbolOfCharacterDown, SymbolOfCharacterLeft, SymbolOfCharacterRight:
				fmt.Printf("【主角: %d 生命值, %s】\n", role.Hp(), role.State())
				return
			}
		}
	}
}

func (g *GameMap) PrintGameMap() {
	fmt.Println("【地圖】")

	for row := 0; row < RowOfGameMap; row++ {
		for col := 0; col < ColOfGameMap; col++ {
			mapObject := g.mapObjects[row][col]

			if mapObject == nil {
				fmt.Printf("%s(%d,%d)\t", "o", col, row)
			} else {
				fmt.Printf("%s(%d,%d)\t", mapObject.String(), col, row)
			}
		}
		fmt.Println()
	}
}

func (g *GameMap) LetRolesTakeTurn() {
	for _, role := range g.roles {
		if role != nil {
			role.TakeTurn(g)
		}
	}
}

type GenerateRandomMapObject func() IMapObject

func (g *GameMap) initializeMapObjects(generateRandomMapObject GenerateRandomMapObject, InitialNumOfMapObjects int) {
	for numOfMapObject := 0; numOfMapObject < InitialNumOfMapObjects; numOfMapObject++ {
		mapObject := generateRandomMapObject()
		row := mapObject.Coordinate().Row()
		col := mapObject.Coordinate().Col()

		g.mapObjects[row][col] = mapObject

		switch mapObject.String() {
		case SymbolOfCharacterUp, SymbolOfCharacterDown, SymbolOfCharacterLeft, SymbolOfCharacterRight, SymbolOfMonster:
			g.roles = append(g.roles, mapObject.(*Role))
		}
	}
}

func (g *GameMap) generateRandomCharacter() IMapObject {
	coordinate, direction := g.generateRandomCoordinateAndDirection()
	return NewRole(coordinate, NewCharacter(direction))
}

func (g *GameMap) generateRandomMonster() IMapObject {
	return NewRole(g.GenerateCoordinateOfRandomEmptyLocation(), NewMonster())
}

func (g *GameMap) generateRandomTreasure() IMapObject {
	return NewTreasure(g.GenerateCoordinateOfRandomEmptyLocation())
}

func (g *GameMap) generateRandomObstacle() IMapObject {
	return NewObstacle(g.GenerateCoordinateOfRandomEmptyLocation())
}

func (g *GameMap) generateRandomCoordinateAndDirection() (*Coordinate, Direction) {
	return g.GenerateCoordinateOfRandomEmptyLocation(), Direction(Rander.Intn(LenOfDirection))
}

func (g *GameMap) GenerateCoordinateOfRandomEmptyLocation() *Coordinate {
	for {
		row := Rander.Intn(RowOfGameMap)
		col := Rander.Intn(ColOfGameMap)

		if g.mapObjects[row][col] == nil {
			coordinate, _ := NewCoordinate(row, col)
			return coordinate
		}
	}
}
