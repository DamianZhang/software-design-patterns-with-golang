package structs

import "fmt"

type Role struct {
	coordinate *Coordinate
	hp         int
	state      *State
	polytype   IRole
}

type IRole interface {
	String() string
	InitialHp() int
	NormalTakeTurn(role *Role, gameMap *GameMap)
	NormalAttacking(role *Role, gameMap *GameMap)
	NormalAttacked(role *Role, damage int, gameMap *GameMap)
	IsFullHp(role *Role) bool
}

func NewRole(coordinate *Coordinate, polytype IRole) *Role {
	r := &Role{
		coordinate: coordinate,
		hp:         polytype.InitialHp(),
		polytype:   polytype,
	}

	r.state = NewState(r, NewNormal())
	return r
}

func (r *Role) String() string {
	return r.polytype.String()
}

func (r *Role) Coordinate() *Coordinate {
	return r.coordinate
}

func (r *Role) SetCoordinate(coordinate *Coordinate) {
	r.coordinate = coordinate
}

func (r *Role) TouchedByRole(role *Role, gameMap *GameMap) {
	fmt.Printf("【%s 觸碰 %s】: 留在原地，不予移動。\n", role.String(), r.String())
}

func (r *Role) Hp() int {
	return r.hp
}

func (r *Role) SetHp(hp int) {
	r.hp = hp
}

func (r *Role) IsDead() bool {
	return r.hp <= 0
}

func (r *Role) State() *State {
	return r.state
}

func (r *Role) EnterState(state IState) {
	r.state.SetState(state)
}

func (r *Role) TakeTurn(gameMap *GameMap) {
	r.state.TakeTurn(gameMap)
}

func (r *Role) Attacking(gameMap *GameMap) {
	r.state.Attacking(gameMap)
}

func (r *Role) Attacked(damage int, gameMap *GameMap) {
	r.state.Attacked(damage, gameMap)
}

func (r *Role) NormalTakeTurn(gameMap *GameMap) {
	r.polytype.NormalTakeTurn(r, gameMap)
}

func (r *Role) NormalAttacking(gameMap *GameMap) {
	r.polytype.NormalAttacking(r, gameMap)
}

func (r *Role) NormalAttacked(damage int, gameMap *GameMap) {
	r.polytype.NormalAttacked(r, damage, gameMap)
}

func (r *Role) IsFullHp() bool {
	return r.polytype.IsFullHp(r)
}

func (r *Role) MoveOneSpaceInOneDirection(direction Direction, gameMap *GameMap) {
	var (
		coordinate *Coordinate
		mapObject  IMapObject
		err        error
		row        int = r.coordinate.Row()
		col        int = r.coordinate.Col()
	)

	switch direction {
	case W:
		coordinate, mapObject, err = r.getCoordinateAndMapObject(row-1, col, gameMap)
	case S:
		coordinate, mapObject, err = r.getCoordinateAndMapObject(row+1, col, gameMap)
	case A:
		coordinate, mapObject, err = r.getCoordinateAndMapObject(row, col-1, gameMap)
	case D:
		coordinate, mapObject, err = r.getCoordinateAndMapObject(row, col+1, gameMap)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	if mapObject == nil {
		gameMap.MoveMapObject(r, coordinate)
	} else {
		mapObject.TouchedByRole(r, gameMap)
	}
}

func (r *Role) getCoordinateAndMapObject(row, col int, gameMap *GameMap) (*Coordinate, IMapObject, error) {
	coordinate, err := NewCoordinate(row, col)
	if err != nil {
		return nil, nil, fmt.Errorf("【%s(%d,%d) 移動失敗】: %s",
			r.String(), r.Coordinate().Col(), r.Coordinate().Row(), err)
	}

	return coordinate, gameMap.GetMapObject(coordinate), nil
}
