package structs

type Database interface {
	GetEmployeeById(id int) (Employee, error)
}
