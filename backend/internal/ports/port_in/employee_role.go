package portin

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
	"thirawoot/in2forest_shop_backend/internal/utils/response"
)

type EmployeeRoleService interface {
	Create(data dto.EmployeeRoleCreate) (*dto.EmployeeRole, error)
	Delete(id uint) (map[string]int64, error)
	Update(id uint, data dto.EmployeeRoleCreate) response.Response
	Find(id uint) response.Response
}
