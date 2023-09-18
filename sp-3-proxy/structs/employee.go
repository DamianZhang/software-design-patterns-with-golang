package structs

type Employee interface {
	Id() int
	Name() string
	Age() int
	IdsOfSubordinates() []int
	Subordinates() []Employee
	AddSubordinate(subordinate Employee)
}
