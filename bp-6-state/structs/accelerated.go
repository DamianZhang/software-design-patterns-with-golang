package structs

import "fmt"

type Accelerated struct{}

func NewAccelerated() *Accelerated {
	return &Accelerated{}
}

func (a *Accelerated) String() string {
	return "加速狀態"
}

func (a *Accelerated) HandleStateAndTakeTurn(state *State, gameMap *GameMap) {
	var (
		role              = state.Role()
		acceleratedRounds = 2
	)

	for acceleratedRound := 0; acceleratedRound < acceleratedRounds; acceleratedRound++ {
		if state.String() == a.String() {
			role.NormalTakeTurn(gameMap)
		}
	}
}

func (a *Accelerated) ChangeStateWhenTimelinessExpired() (timeliness int, changeState IState) {
	return 3, NewNormal()
}

func (a *Accelerated) Attacking(state *State, gameMap *GameMap) {
	state.Role().NormalAttacking(gameMap)
}

func (a *Accelerated) Attacked(state *State, damage int, gameMap *GameMap) {
	var (
		role              = state.Role()
		informationOfRole = fmt.Sprintf("%s(%d,%d)", role.String(), role.Coordinate().Col(), role.Coordinate().Row())
	)

	role.SetHp(role.Hp() - damage)
	fmt.Printf("%s 遭受 %d 點攻擊，剩餘 %d 點生命值\n", informationOfRole, damage, role.Hp())

	if role.IsDead() {
		gameMap.RemoveMapObject(role.Coordinate())
		gameMap.RemoveRole(role)
		fmt.Printf("%s 死亡，被移除遊戲地圖\n", informationOfRole)
	} else {
		role.EnterState(NewNormal())
	}
}
