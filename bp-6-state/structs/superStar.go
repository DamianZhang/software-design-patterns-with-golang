package structs

type SuperStar struct{}

func NewSuperStar() *SuperStar {
	return &SuperStar{}
}

func (s *SuperStar) TouchEffect(role *Role) {
	role.EnterState(NewInvincible())
}
