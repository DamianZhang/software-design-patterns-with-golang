package structs

type Treasure struct {
	coordinate      *Coordinate
	treasureContent ITreasureContent
}

type ITreasureContent interface {
	TouchEffect(role *Role)
}

func NewTreasure(coordinate *Coordinate) *Treasure {
	return &Treasure{
		coordinate:      coordinate,
		treasureContent: generateRandomTreasureContent(),
	}
}

func (t *Treasure) String() string {
	return SymbolOfTreasure
}

func (t *Treasure) Coordinate() *Coordinate {
	return t.coordinate
}

func (t *Treasure) SetCoordinate(coordinate *Coordinate) {
	t.coordinate = coordinate
}

func (t *Treasure) TouchedByRole(role *Role, gameMap *GameMap) {
	gameMap.MoveMapObject(role, t.Coordinate())
	t.treasureContent.TouchEffect(role)
}

func generateRandomTreasureContent() ITreasureContent {
	number := Rander.Intn(100)
	switch {
	case number >= 90:
		return NewSuperStar()
	case number >= 65:
		return NewPoison()
	case number >= 45:
		return NewAcceleratingPotion()
	case number >= 30:
		return NewHealingPotion()
	case number >= 20:
		return NewDevilFruit()
	case number >= 10:
		return NewKingsRock()
	case number >= 0:
		return NewDokodemoDoor()
	default:
		return NewSuperStar()
	}
}
