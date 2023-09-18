package structs

type Normal struct{}

func NewNormal() *Normal {
	return &Normal{}
}

func (n *Normal) String() string {
	return "正常狀態"
}

func (n *Normal) HandleStateAndTakeTurn(state *State, gameMap *GameMap) {
	state.Role().NormalTakeTurn(gameMap)
}

func (n *Normal) ChangeStateWhenTimelinessExpired() (timeliness int, changeState IState) {
	return -1, NewNormal()
}

func (n *Normal) Attacking(state *State, gameMap *GameMap) {
	state.Role().NormalAttacking(gameMap)
}

func (n *Normal) Attacked(state *State, damage int, gameMap *GameMap) {
	state.Role().NormalAttacked(damage, gameMap)
}
