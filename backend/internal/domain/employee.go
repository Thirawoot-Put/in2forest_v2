package domain

import "time"

type Employee struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Email    string
	Password string
	Name     string
	Mobile   string

	EmployeeRoleID uint
}
