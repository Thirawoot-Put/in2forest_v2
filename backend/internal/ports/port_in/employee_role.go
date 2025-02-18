package portin

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
)

type EmployeeRoleApp interface {
	Create(data dto.EmployeeRoleCreate) (*dto.EmployeeRole, error)
	Delete(id uint) (map[string]int64, error)
	Update(id uint, data dto.EmployeeRoleCreate) (map[string]int64, error)
	Find(id uint) (*dto.EmployeeRole, error)
	FindAll() ([]dto.EmployeeRole, error)
	FindByRole(role string) (*dto.EmployeeRole, error)
}
