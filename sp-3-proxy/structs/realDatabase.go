package structs

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RealDatabase struct {
	dataSourceFileName string
}

func NewRealDatabase(dataSourceFileName string) *RealDatabase {
	return &RealDatabase{
		dataSourceFileName: dataSourceFileName,
	}
}

func (r *RealDatabase) GetEmployeeById(id int) (Employee, error) {
	bytesOfData, err := os.ReadFile(r.dataSourceFileName)
	if err != nil {
		return nil, fmt.Errorf("data source load failed: %s", err)
	}

	var (
		data      = string(bytesOfData)
		employees = strings.Split(data, "\n")
	)

	if id < 1 || id > len(employees)-1 {
		return nil, errors.New("illegal employee ID")
	}

	var (
		employee       = employees[id]
		infoOfEmployee = strings.Split(employee, " ")
	)

	idOfEmployee, err := strconv.Atoi(infoOfEmployee[0])
	if err != nil {
		return nil, fmt.Errorf("illegal employee ID: %s", err)
	}

	ageOfEmployee, err := strconv.Atoi(infoOfEmployee[2])
	if err != nil {
		return nil, fmt.Errorf("illegal employee age: %s", err)
	}

	var (
		nameOfEmployee                       = infoOfEmployee[1]
		stringsOfIdsOfSubordinatesOfEmployee = strings.Split(infoOfEmployee[3], ",")
		intsOfIdsOfSubordinatesOfEmployee    = make([]int, 0)
	)

	if len(stringsOfIdsOfSubordinatesOfEmployee) != 0 && stringsOfIdsOfSubordinatesOfEmployee[0] != "" {
		for _, stringOfIdOfSubordinateOfEmployee := range stringsOfIdsOfSubordinatesOfEmployee {
			intOfIdOfSubordinateOfEmployee, err := strconv.Atoi(stringOfIdOfSubordinateOfEmployee)
			if err != nil {
				return nil, fmt.Errorf("illegal subordinate ID: %s", err)
			}

			intsOfIdsOfSubordinatesOfEmployee = append(intsOfIdsOfSubordinatesOfEmployee, intOfIdOfSubordinateOfEmployee)
		}
	}

	return NewRealEmployee(idOfEmployee, nameOfEmployee, ageOfEmployee, intsOfIdsOfSubordinatesOfEmployee), nil
}
