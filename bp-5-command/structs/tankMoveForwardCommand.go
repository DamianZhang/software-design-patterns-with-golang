package structs

type TankMoveForwardCommand struct {
	tank *Tank
}

func NewTankMoveForwardCommand(tank *Tank) *TankMoveForwardCommand {
	return &TankMoveForwardCommand{tank: tank}
}

func (t *TankMoveForwardCommand) Execute() {
	t.tank.MoveForward()
}

func (t *TankMoveForwardCommand) Undo() {
	t.tank.MoveBackward()
}

func (t *TankMoveForwardCommand) String() string {
	return "MoveTankForward"
}
