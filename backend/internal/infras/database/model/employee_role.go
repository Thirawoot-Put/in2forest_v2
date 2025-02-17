package model

import "gorm.io/gorm"

type EmployeeRole struct {
	gorm.Model

	Role string `gorm:"unique;not null;size:100"`

	Employees []Employee
}
