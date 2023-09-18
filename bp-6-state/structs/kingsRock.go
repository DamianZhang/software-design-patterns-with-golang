package structs

type KingsRock struct{}

func NewKingsRock() *KingsRock {
	return &KingsRock{}
}

func (k *KingsRock) TouchEffect(role *Role) {
	role.EnterState(NewStockpile())
}
