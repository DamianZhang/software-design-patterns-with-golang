package structs

type AcceleratingPotion struct{}

func NewAcceleratingPotion() *AcceleratingPotion {
	return &AcceleratingPotion{}
}

func (a *AcceleratingPotion) TouchEffect(role *Role) {
	role.EnterState(NewAccelerated())
}
