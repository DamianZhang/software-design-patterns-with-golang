package structs

import "fmt"

type VirtualRealEmployeeProxy struct {
	id                int
	name              string
	age               int
	idsOfSubordinates []int
	realEmployee      *RealEmployee
	db                Database
}

func NewVirtualRealEmployeeProxy(id int, name string, age int, idsOfSubordinates []int, db Database) *VirtualRealEmployeeProxy {
	return &VirtualRealEmployeeProxy{
		id:                id,
		name:              name,
		age:               age,
		idsOfSubordinates: idsOfSubordinates,
		realEmployee:      nil,
		db:                db,
	}
}

func (v *VirtualRealEmployeeProxy) lazyInitialization() {
	if v.realEmployee == nil {
		fmt.Println("【觸發延遲載入】")
		v.realEmployee = NewRealEmployee(v.id, v.name, v.age, v.idsOfSubordinates)
	}
}

func (v *VirtualRealEmployeeProxy) lazyLoadSubordinates() {
	v.lazyInitialization()

	if len(v.realEmployee.Subordinates()) == 0 {
		for _, idOfSubordinate := range v.realEmployee.IdsOfSubordinates() {
			subordinate, err := v.db.GetEmployeeById(idOfSubordinate)
			if err != nil {
				fmt.Println("virtual proxy lazy load subordinates failed:", err)
				return
			}

			v.realEmployee.AddSubordinate(subordinate)
		}
	}
}

func (v *VirtualRealEmployeeProxy) Id() int {
	v.lazyInitialization()
	return v.realEmployee.Id()
}

func (v *VirtualRealEmployeeProxy) Name() string {
	v.lazyInitialization()
	return v.realEmployee.Name()
}

func (v *VirtualRealEmployeeProxy) Age() int {
	v.lazyInitialization()
	return v.realEmployee.Age()
}

func (v *VirtualRealEmployeeProxy) IdsOfSubordinates() []int {
	v.lazyInitialization()
	return v.realEmployee.IdsOfSubordinates()
}

func (v *VirtualRealEmployeeProxy) Subordinates() []Employee {
	v.lazyLoadSubordinates()
	return v.realEmployee.Subordinates()
}

func (v *VirtualRealEmployeeProxy) AddSubordinate(subordinate Employee) {
	v.lazyLoadSubordinates()
	v.realEmployee.AddSubordinate(subordinate)
}
