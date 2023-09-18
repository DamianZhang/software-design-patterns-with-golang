package structs

type DokodemoDoor struct{}

func NewDokodemoDoor() *DokodemoDoor {
	return &DokodemoDoor{}
}

func (d *DokodemoDoor) TouchEffect(role *Role) {
	role.EnterState(NewTeleport())
}
