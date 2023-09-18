package structs

import "fmt"

type Healing struct{}

func NewHealing() *Healing {
	return &Healing{}
}

func (h *Healing) String() string {
	return "恢復狀態"
}

func (h *Healing) HandleStateAndTakeTurn(state *State, gameMap *GameMap) {
	var (
		role              = state.Role()
		informationOfRole = fmt.Sprintf("%s(%d,%d)", role.String(), role.Coordinate().Col(), role.Coordinate().Row())
		recoveryHp        = 30
	)

	if role.IsFullHp() {
		role.EnterState(NewNormal())
	} else {
		role.SetHp(role.Hp() + recoveryHp)
		fmt.Printf("%s 恢復 %d 點生命值，剩餘 %d 點生命值\n", informationOfRole, recoveryHp, role.Hp())
	}

	role.NormalTakeTurn(gameMap)
}

func (h *Healing) ChangeStateWhenTimelinessExpired() (timeliness int, changeState IState) {
	return 5, NewNormal()
}

func (h *Healing) Attacking(state *State, gameMap *GameMap) {
	state.Role().NormalAttacking(gameMap)
}

func (h *Healing) Attacked(state *State, damage int, gameMap *GameMap) {
	state.Role().NormalAttacked(damage, gameMap)
}
