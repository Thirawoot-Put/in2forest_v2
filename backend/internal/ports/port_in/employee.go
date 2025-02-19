package portin

import "thirawoot/in2forest_shop_backend/internal/dto"

type EmployeeApp interface {
	Create(data dto.Employee) (*dto.Employee, error)
	FindByEmail(email string) (*dto.Employee, error)
}
