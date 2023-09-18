package structs

import (
	"errors"
	"fmt"
	"os"
)

const (
	SP_3_PROXY_PASSWORD = "setpassword"
)

type ProtectionRealDatabaseProxy struct {
	realDatabase *RealDatabase
}

func NewProtectionRealDatabaseProxy(dataSourceFileName string) *ProtectionRealDatabaseProxy {
	return &ProtectionRealDatabaseProxy{
		realDatabase: NewRealDatabase(dataSourceFileName),
	}
}

func (p *ProtectionRealDatabaseProxy) GetEmployeeById(id int) (Employee, error) {
	if os.Getenv("SP_3_PROXY_PASSWORD") != SP_3_PROXY_PASSWORD {
		return nil, errors.New("403 Forbidden")
	}

	employee, err := p.realDatabase.GetEmployeeById(id)
	if err != nil {
		return nil, fmt.Errorf("protection proxy GetEmployeeById failed: %s", err)
	}

	return NewVirtualRealEmployeeProxy(employee.Id(), employee.Name(), employee.Age(), employee.IdsOfSubordinates(), p), nil
}
