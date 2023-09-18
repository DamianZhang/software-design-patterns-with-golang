package structs

type Teleport struct{}

func NewTeleport() *Teleport {
	return &Teleport{}
}

func (t *Teleport) String() string {
	return "瞬身狀態"
}

func (t *Teleport) HandleStateAndTakeTurn(state *State, gameMap *GameMap) {
	role := state.Role()

	gameMap.MoveMapObject(role, gameMap.GenerateCoordinateOfRandomEmptyLocation())
	role.NormalTakeTurn(gameMap)
}

func (t *Teleport) ChangeStateWhenTimelinessExpired() (timeliness int, changeState IState) {
	return 1, NewNormal()
}

func (t *Teleport) Attacking(state *State, gameMap *GameMap) {
	state.Role().NormalAttacking(gameMap)
}

func (t *Teleport) Attacked(state *State, damage int, gameMap *GameMap) {
	state.Role().NormalAttacked(damage, gameMap)
}
