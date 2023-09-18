package structs

import (
	"fmt"
	"regexp"
	"strconv"
)

type Client struct {
	db Database
}

func NewClient(db Database) *Client {
	return &Client{
		db: db,
	}
}

func (c *Client) Start() {
	fmt.Println("employee database is starting...")

	var stringOfEmployeeId string

	for {
		fmt.Println("please input a employee ID:")
		fmt.Scanf("%s", &stringOfEmployeeId)

		if stringOfEmployeeId == "exit" {
			fmt.Println("Database is closed...")
			break
		}

		if IsLegalStringOfEmployeeId(stringOfEmployeeId) {
			employeeId, _ := strconv.Atoi(stringOfEmployeeId)
			employee, err := c.db.GetEmployeeById(employeeId)
			if err != nil {
				fmt.Println("client GetEmployeeById failed:", err)
				continue
			}

			fmt.Printf("subordinates of employee %s:\n", stringOfEmployeeId)
			for _, subordinate := range employee.Subordinates() {
				fmt.Printf("%d: %s\n", subordinate.Id(), subordinate.Name())
			}
		} else {
			fmt.Println("please input a LEGAL employee ID!!!")
		}
	}
}

func IsLegalStringOfEmployeeId(stringOfEmployeeId string) bool {
	match, _ := regexp.MatchString("\\b[0-9]{1,9}\\b", stringOfEmployeeId)
	return match
}
