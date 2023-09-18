package structs

import "fmt"

type Invincible struct{}

func NewInvincible() *Invincible {
	return &Invincible{}
}

func (i *Invincible) String() string {
	return "無敵狀態"
}

func (i *Invincible) HandleStateAndTakeTurn(state *State, gameMap *GameMap) {
	state.Role().NormalTakeTurn(gameMap)
}

func (i *Invincible) ChangeStateWhenTimelinessExpired() (timeliness int, changeState IState) {
	return 2, NewNormal()
}

func (i *Invincible) Attacking(state *State, gameMap *GameMap) {
	state.Role().NormalAttacking(gameMap)
}

func (i *Invincible) Attacked(state *State, damage int, gameMap *GameMap) {
	fmt.Println("【攻擊無效】: 無敵狀態")
}
