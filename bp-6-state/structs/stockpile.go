package structs

import "fmt"

type Stockpile struct{}

func NewStockpile() *Stockpile {
	return &Stockpile{}
}

func (s *Stockpile) String() string {
	return "蓄力狀態"
}

func (s *Stockpile) HandleStateAndTakeTurn(state *State, gameMap *GameMap) {
	state.Role().NormalTakeTurn(gameMap)
}

func (s *Stockpile) ChangeStateWhenTimelinessExpired() (timeliness int, changeState IState) {
	return 2, NewErupting()
}

func (s *Stockpile) Attacking(state *State, gameMap *GameMap) {
	state.Role().NormalAttacking(gameMap)
}

func (s *Stockpile) Attacked(state *State, damage int, gameMap *GameMap) {
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
