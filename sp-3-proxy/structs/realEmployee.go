package structs

type RealEmployee struct {
	id                int
	name              string
	age               int
	idsOfSubordinates []int
	subordinates      []Employee
}

func NewRealEmployee(id int, name string, age int, idsOfSubordinates []int) *RealEmployee {
	return &RealEmployee{
		id:                id,
		name:              name,
		age:               age,
		idsOfSubordinates: idsOfSubordinates,
		subordinates:      make([]Employee, 0),
	}
}

func (r *RealEmployee) Id() int {
	return r.id
}

func (r *RealEmployee) Name() string {
	return r.name
}

func (r *RealEmployee) Age() int {
	return r.age
}

func (r *RealEmployee) IdsOfSubordinates() []int {
	return r.idsOfSubordinates
}

func (r *RealEmployee) Subordinates() []Employee {
	return r.subordinates
}

func (r *RealEmployee) AddSubordinate(subordinate Employee) {
	r.subordinates = append(r.subordinates, subordinate)
}
