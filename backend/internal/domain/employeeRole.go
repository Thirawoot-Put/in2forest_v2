package domain

import "time"

type EmployeeRole struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Role string
}
