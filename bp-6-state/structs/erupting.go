package structs

type Erupting struct{}

func NewErupting() *Erupting {
	return &Erupting{}
}

func (s *Erupting) String() string {
	return "爆發狀態"
}

func (e *Erupting) HandleStateAndTakeTurn(state *State, gameMap *GameMap) {
	state.Role().NormalTakeTurn(gameMap)
}

func (e *Erupting) ChangeStateWhenTimelinessExpired() (timeliness int, changeState IState) {
	return 3, NewTeleport()
}

func (e *Erupting) Attacking(state *State, gameMap *GameMap) {
	var (
		role           = state.Role()
		roles          = gameMap.Roles()
		eruptingDamage = 100
	)

	for _, r := range roles {
		if r != nil {
			if r.String() == role.String() {
				continue
			}

			r.Attacked(eruptingDamage, gameMap)
		}
	}
}

func (e *Erupting) Attacked(state *State, damage int, gameMap *GameMap) {
	state.Role().NormalAttacked(damage, gameMap)
}
