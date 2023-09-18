package structs

import "fmt"

type State struct {
	timeliness int
	role       *Role
	polytype   IState
}

type IState interface {
	String() string
	HandleStateAndTakeTurn(state *State, gameMap *GameMap)
	ChangeStateWhenTimelinessExpired() (timeliness int, changeState IState)
	Attacking(state *State, gameMap *GameMap)
	Attacked(state *State, damage int, gameMap *GameMap)
}

func NewState(role *Role, polytype IState) *State {
	return &State{
		timeliness: 0,
		role:       role,
		polytype:   polytype,
	}
}

func (s *State) SetState(polytype IState) {
	s.ExitState()
	s.polytype = polytype
	s.EnterState()

	fmt.Printf("%s(%d,%d) 進入 %s\n", s.role.String(), s.role.Coordinate().Col(), s.role.Coordinate().Row(), s.String())
}

func (s *State) EnterState() {
	s.timeliness = 0
}

func (s *State) ExitState() {
	s.timeliness = 0
}

func (s *State) String() string {
	return s.polytype.String()
}

func (s *State) TakeTurn(gameMap *GameMap) {
	s.timeliness++

	s.polytype.HandleStateAndTakeTurn(s, gameMap)

	timeliness, changeState := s.polytype.ChangeStateWhenTimelinessExpired()
	s.changeStateWhenTimelinessExpired(timeliness, changeState)
}

func (s *State) changeStateWhenTimelinessExpired(timeliness int, changeState IState) {
	if s.timeliness == timeliness {
		s.SetState(changeState)
	}
}

func (s *State) Attacking(gameMap *GameMap) {
	s.polytype.Attacking(s, gameMap)
}

func (s *State) Attacked(damage int, gameMap *GameMap) {
	s.polytype.Attacked(s, damage, gameMap)
}

func (s *State) Role() *Role {
	return s.role
}
