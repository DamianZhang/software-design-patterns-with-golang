package structs

type Water struct{}

func NewWater() *Water {
	return &Water{}
}

func (w Water) String() string {
	return "W"
}

func (w *Water) Hp() int {
	return 0
}

func (w *Water) SetHp(hp int) {}

func (w *Water) IsDead() bool {
	return false
}
