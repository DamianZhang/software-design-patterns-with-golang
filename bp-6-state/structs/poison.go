package structs

type Poison struct{}

func NewPoison() *Poison {
	return &Poison{}
}

func (p *Poison) TouchEffect(role *Role) {
	role.EnterState(NewPoisoned())
}
