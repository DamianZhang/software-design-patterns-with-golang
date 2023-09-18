package structs

type DevilFruit struct{}

func NewDevilFruit() *DevilFruit {
	return &DevilFruit{}
}

func (d *DevilFruit) TouchEffect(role *Role) {
	role.EnterState(NewOrderless())
}
