package structs

type TankMoveBackwardCommand struct {
	tank *Tank
}

func NewTankMoveBackwardCommand(tank *Tank) *TankMoveBackwardCommand {
	return &TankMoveBackwardCommand{tank: tank}
}

func (t *TankMoveBackwardCommand) Execute() {
	t.tank.MoveBackward()
}

func (t *TankMoveBackwardCommand) Undo() {
	t.tank.MoveForward()
}

func (t *TankMoveBackwardCommand) String() string {
	return "MoveTankBackward"
}
