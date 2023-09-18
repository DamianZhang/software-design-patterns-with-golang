package structs

type HealingPotion struct{}

func NewHealingPotion() *HealingPotion {
	return &HealingPotion{}
}

func (h *HealingPotion) TouchEffect(role *Role) {
	role.EnterState(NewHealing())
}
