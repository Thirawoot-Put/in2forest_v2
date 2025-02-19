package mapper

import (
	"thirawoot/in2forest_shop_backend/internal/domain"
	"thirawoot/in2forest_shop_backend/internal/dto"
	"thirawoot/in2forest_shop_backend/internal/infras/database/model"
)

func ToEmployeeRoleDto(domain domain.EmployeeRole) dto.EmployeeRole {
	return dto.EmployeeRole{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func ModelToEmployeeRoleDomain(m model.EmployeeRole) domain.EmployeeRole {
	return domain.EmployeeRole{
		ID:   m.ID,
		Name: m.Name,
	}
}

func DomainToEmployeeRoleModel(d domain.EmployeeRole) model.EmployeeRole {
	return model.EmployeeRole{
		Name: d.Name,
	}
}
