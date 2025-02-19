package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model

	Email    string `gorm:"unique;not null;size(100)"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Mobile   string `gorm:"not null"`

	RoleID uint
	Role   EmployeeRole `gorm:"foreignKey:RoleID"`
}
