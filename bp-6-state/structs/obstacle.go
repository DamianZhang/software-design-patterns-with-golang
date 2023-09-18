package structs

import "fmt"

type Obstacle struct {
	coordinate *Coordinate
}

func NewObstacle(coordinate *Coordinate) *Obstacle {
	return &Obstacle{
		coordinate: coordinate,
	}
}

func (o *Obstacle) String() string {
	return SymbolOfObstacle
}

func (o *Obstacle) Coordinate() *Coordinate {
	return o.coordinate
}

func (o *Obstacle) SetCoordinate(coordinate *Coordinate) {
	o.coordinate = coordinate
}

func (o *Obstacle) TouchedByRole(role *Role, gameMap *GameMap) {
	fmt.Printf("【%s(%d,%d) 觸碰 %s】: 留在原地，不予移動。\n",
		role.String(), role.Coordinate().Col(), role.Coordinate().Row(), o.String())
}
