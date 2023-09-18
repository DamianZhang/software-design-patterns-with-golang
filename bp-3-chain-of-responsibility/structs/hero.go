package structs

const (
	INITIAL_HP int = 30
)

type Hero struct {
	hp int
}

func NewHero() *Hero {
	return &Hero{hp: INITIAL_HP}
}

func (h Hero) String() string {
	return "H"
}

func (h *Hero) Hp() int {
	return h.hp
}

func (h *Hero) SetHp(hp int) {
	h.hp = hp
}

func (h *Hero) IsDead() bool {
	return h.hp <= 0
}
