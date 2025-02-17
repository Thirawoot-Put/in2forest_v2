package portin

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
	"thirawoot/in2forest_shop_backend/internal/utils/response"
)

type EmployeeRoleService interface {
	Create(data dto.EmployeeRoleCreate) response.Response
}
