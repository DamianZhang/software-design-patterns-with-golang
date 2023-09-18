package structs

import "fmt"

type Poisoned struct{}

func NewPoisoned() *Poisoned {
	return &Poisoned{}
}

func (p *Poisoned) String() string {
	return "中毒狀態"
}

func (p *Poisoned) HandleStateAndTakeTurn(state *State, gameMap *GameMap) {
	var (
		role              = state.Role()
		informationOfRole = fmt.Sprintf("%s(%d,%d)", role.String(), role.Coordinate().Col(), role.Coordinate().Row())
		poisonedDamage    = 15
	)

	role.SetHp(role.Hp() - poisonedDamage)
	fmt.Printf("%s 遭受 %d 點攻擊，剩餘 %d 點生命值\n", informationOfRole, poisonedDamage, role.Hp())

	if role.IsDead() {
		gameMap.RemoveMapObject(role.Coordinate())
		gameMap.RemoveRole(role)
		fmt.Printf("%s 死亡，被移除遊戲地圖\n", informationOfRole)
	} else {
		role.NormalTakeTurn(gameMap)
	}
}

func (p *Poisoned) ChangeStateWhenTimelinessExpired() (timeliness int, changeState IState) {
	return 3, NewNormal()
}

func (p *Poisoned) Attacking(state *State, gameMap *GameMap) {
	state.Role().NormalAttacking(gameMap)
}

func (p *Poisoned) Attacked(state *State, damage int, gameMap *GameMap) {
	state.Role().NormalAttacked(damage, gameMap)
}
