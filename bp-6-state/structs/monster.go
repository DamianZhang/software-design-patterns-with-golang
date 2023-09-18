package structs

import "fmt"

const (
	InitialHpOfMonster int = 1
)

type Monster struct{}

func NewMonster() *Monster {
	return &Monster{}
}

func (m *Monster) String() string {
	return SymbolOfMonster
}

func (m *Monster) InitialHp() int {
	return InitialHpOfMonster
}

func (m *Monster) NormalTakeTurn(role *Role, gameMap *GameMap) {
	if character := m.getSurroundedByCharacter(role, gameMap); character != nil {
		role.Attacking(gameMap)
	} else {
		direction := Direction(Rander.Intn(LenOfDirection))
		role.MoveOneSpaceInOneDirection(direction, gameMap)
	}
}

func (m *Monster) getSurroundedByCharacter(role *Role, gameMap *GameMap) IMapObject {
	surroundedByMapObjects := m.getSurroundedByMapObjects(role, gameMap)

	for _, mapObject := range surroundedByMapObjects {
		if mapObject != nil {
			switch mapObject.String() {
			case SymbolOfCharacterUp, SymbolOfCharacterDown, SymbolOfCharacterLeft, SymbolOfCharacterRight:
				return mapObject
			}
		}
	}

	return nil
}

func (m *Monster) getSurroundedByMapObjects(role *Role, gameMap *GameMap) []IMapObject {
	var (
		row                    = role.Coordinate().Row()
		col                    = role.Coordinate().Col()
		surroundedByMapObjects = make([]IMapObject, 0)
	)

	surroundedByMapObjects = append(surroundedByMapObjects,
		m.getMapObject(row-1, col, gameMap), m.getMapObject(row+1, col, gameMap),
		m.getMapObject(row, col-1, gameMap), m.getMapObject(row, col+1, gameMap),
	)

	return surroundedByMapObjects
}

func (m *Monster) getMapObject(row, col int, gameMap *GameMap) IMapObject {
	coordinate, err := NewCoordinate(row, col)
	if err != nil {
		return nil
	}

	return gameMap.GetMapObject(coordinate)
}

func (m *Monster) NormalAttacking(role *Role, gameMap *GameMap) {
	if character := m.getSurroundedByCharacter(role, gameMap); character != nil {
		character.(*Role).Attacked(50, gameMap)
	}
}

func (m *Monster) NormalAttacked(role *Role, damage int, gameMap *GameMap) {
	informationOfRole := fmt.Sprintf("%s(%d,%d)", role.String(), role.Coordinate().Col(), role.Coordinate().Row())

	role.SetHp(role.Hp() - damage)
	fmt.Printf("%s 遭受 %d 點攻擊，剩餘 %d 點生命值\n", informationOfRole, damage, role.Hp())

	if role.IsDead() {
		gameMap.RemoveMapObject(role.Coordinate())
		gameMap.RemoveRole(role)
		fmt.Printf("%s 死亡，被移除遊戲地圖\n", informationOfRole)
	}
}

func (m *Monster) IsFullHp(role *Role) bool {
	return role.Hp() == InitialHpOfMonster
}
