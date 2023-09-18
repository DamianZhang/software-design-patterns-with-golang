package structs

import (
	"fmt"
)

const (
	OnlyMoveUpDown    int = 0
	OnlyMoveLeftRight int = 1
)

type Orderless struct{}

func NewOrderless() *Orderless {
	return &Orderless{}
}

func (o *Orderless) String() string {
	return "混亂狀態"
}

func (o *Orderless) HandleStateAndTakeTurn(state *State, gameMap *GameMap) {
	var (
		role              = state.Role()
		informationOfRole = fmt.Sprintf("%s(%d,%d)", role.String(), role.Coordinate().Col(), role.Coordinate().Row())
		orderlessType     = Rander.Intn(2)
	)

	switch orderlessType {
	case OnlyMoveUpDown:
		promptWord := fmt.Sprintf("%s 請選擇 (W) 上 (S) 下:", informationOfRole)
		o.onlyMoveTwoDirections(promptWord, W, S, role, gameMap)
	case OnlyMoveLeftRight:
		promptWord := fmt.Sprintf("%s 請選擇 (A) 左 (D) 右:", informationOfRole)
		o.onlyMoveTwoDirections(promptWord, A, D, role, gameMap)
	}
}

func (o *Orderless) onlyMoveTwoDirections(promptWord string, directionA, directionB Direction, role *Role, gameMap *GameMap) {
	var (
		alphabet             string
		alphabetOfDirectionA      = directionA.Alphabet()
		alphabetOfDirectionB      = directionB.Alphabet()
		exit                 bool = false
	)

	for !exit {
		fmt.Println(promptWord)
		fmt.Scanf("%s", &alphabet)

		switch alphabet {
		case alphabetOfDirectionA:
			role.MoveOneSpaceInOneDirection(directionA, gameMap)
			exit = true
		case alphabetOfDirectionB:
			role.MoveOneSpaceInOneDirection(directionB, gameMap)
			exit = true
		default:
			fmt.Printf("請輸入 %s, %s 以執行程序\n", alphabetOfDirectionA, alphabetOfDirectionB)
		}
	}
}

func (o *Orderless) ChangeStateWhenTimelinessExpired() (timeliness int, changeState IState) {
	return 3, NewNormal()
}

func (*Orderless) Attacking(state *State, gameMap *GameMap) {
	state.Role().NormalAttacking(gameMap)
}

func (o *Orderless) Attacked(state *State, damage int, gameMap *GameMap) {
	state.Role().NormalAttacked(damage, gameMap)
}
