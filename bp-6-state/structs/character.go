package structs

import "fmt"

const (
	InitialHpOfCharacter int = 300
	AttackingVertical    int = 1
	AttackingHorizontal  int = 2
)

type Character struct {
	direction Direction
}

func NewCharacter(direction Direction) *Character {
	return &Character{
		direction: direction,
	}
}

func (c *Character) String() string {
	return c.direction.String()
}

func (c *Character) InitialHp() int {
	return InitialHpOfCharacter
}

func (c *Character) NormalTakeTurn(role *Role, gameMap *GameMap) {
	var (
		input string
		exit  bool = false
		row        = role.Coordinate().Row()
		col        = role.Coordinate().Col()
	)

	for !exit {
		fmt.Printf("現在位置(%d,%d) (1) 移動 (2) 攻擊:\n", col, row)
		fmt.Scanf("%s", &input)

		switch input {
		case "1":
			movingDirection := c.scanDirection("移動 (W) 上 (S) 下 (A) 左 (D) 右:")
			facingDirection := c.scanDirection("面向 (W) 上 (S) 下 (A) 左 (D) 右:")
			c.direction = facingDirection
			role.MoveOneSpaceInOneDirection(movingDirection, gameMap)
			exit = true
		case "2":
			role.Attacking(gameMap)
			exit = true
		default:
			fmt.Println("請輸入數字 1-2 以執行程序")
		}
	}
}

func (c *Character) scanDirection(promptWord string) Direction {
	var (
		alphabet    string
		alphabetOfW = W.Alphabet()
		alphabetOfS = S.Alphabet()
		alphabetOfA = A.Alphabet()
		alphabetOfD = D.Alphabet()
	)

	for {
		fmt.Println(promptWord)
		fmt.Scanf("%s", &alphabet)

		switch alphabet {
		case alphabetOfW, alphabetOfS, alphabetOfA, alphabetOfD:
			direction, _ := ConvertAlphabetToDirection(alphabet)
			return direction
		default:
			fmt.Println("請輸入 W, S, A, D 以執行程序")
		}
	}
}

func (c *Character) NormalAttacking(role *Role, gameMap *GameMap) {
	var (
		direction = c.direction
		row       = role.Coordinate().Row()
		col       = role.Coordinate().Col()
	)

	switch direction {
	case W:
		c.attackingMapObjects(AttackingVertical, row-1, col,
			func(num int) bool { return num >= 0 },
			func(num int) int { return num - 1 },
			gameMap,
		)
	case S:
		c.attackingMapObjects(AttackingVertical, row+1, col,
			func(num int) bool { return num < RowOfGameMap },
			func(num int) int { return num + 1 },
			gameMap,
		)
	case A:
		c.attackingMapObjects(AttackingHorizontal, col-1, row,
			func(num int) bool { return num >= 0 },
			func(num int) int { return num - 1 },
			gameMap,
		)
	case D:
		c.attackingMapObjects(AttackingHorizontal, col+1, row,
			func(num int) bool { return num < ColOfGameMap },
			func(num int) int { return num + 1 },
			gameMap,
		)
	}
}

type LoopCondition func(num int) bool
type AfterOneLoop func(num int) int

func (c *Character) attackingMapObjects(attackingType, startNum, fixedNum int,
	loopCondition LoopCondition, afterOneLoop AfterOneLoop, gameMap *GameMap) {
	var mapObject IMapObject

	for num := startNum; loopCondition(num); {
		switch attackingType {
		case AttackingVertical:
			coordinate, _ := NewCoordinate(num, fixedNum)
			mapObject = gameMap.GetMapObject(coordinate)
		case AttackingHorizontal:
			coordinate, _ := NewCoordinate(fixedNum, num)
			mapObject = gameMap.GetMapObject(coordinate)
		}

		if mapObject != nil {
			switch mapObject.String() {
			case SymbolOfMonster:
				mapObject.(*Role).Attacked(mapObject.(*Role).Hp(), gameMap)
			case SymbolOfObstacle:
				fmt.Println("【無效攻擊】: 無法攻擊障礙物")
				return
			}
		}

		num = afterOneLoop(num)
	}
}

func (c *Character) NormalAttacked(role *Role, damage int, gameMap *GameMap) {
	informationOfRole := fmt.Sprintf("%s(%d,%d)", role.String(), role.Coordinate().Col(), role.Coordinate().Row())

	role.SetHp(role.Hp() - damage)
	fmt.Printf("%s 遭受 %d 點攻擊，剩餘 %d 點生命值\n", informationOfRole, damage, role.Hp())

	if role.IsDead() {
		gameMap.RemoveMapObject(role.Coordinate())
		gameMap.RemoveRole(role)
		fmt.Printf("%s 死亡，被移除遊戲地圖\n", informationOfRole)
	} else {
		role.EnterState(NewInvincible())
	}
}

func (c *Character) IsFullHp(role *Role) bool {
	return role.Hp() == InitialHpOfCharacter
}
