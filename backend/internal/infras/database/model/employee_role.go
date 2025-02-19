package model

import "gorm.io/gorm"

type EmployeeRole struct {
	gorm.Model

	Name string `gorm:"unique;not null;size:100"`
}
