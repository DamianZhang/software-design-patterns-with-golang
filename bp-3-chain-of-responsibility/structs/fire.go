package structs

type Fire struct{}

func NewFire() *Fire {
	return &Fire{}
}

func (f Fire) String() string {
	return "F"
}

func (f *Fire) Hp() int {
	return 0
}

func (f *Fire) SetHp(hp int) {}

func (f *Fire) IsDead() bool {
	return false
}
