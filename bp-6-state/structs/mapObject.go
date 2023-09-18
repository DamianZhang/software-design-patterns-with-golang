package structs

type IMapObject interface {
	String() string
	Coordinate() *Coordinate
	SetCoordinate(coordinate *Coordinate)
	TouchedByRole(role *Role, gameMap *GameMap)
}
