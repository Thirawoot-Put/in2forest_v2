package mapper

import (
	"thirawoot/in2forest_shop_backend/internal/domain"
	"thirawoot/in2forest_shop_backend/internal/dto"
)

func ToEmployeeRoleDto(domain domain.EmployeeRole) dto.EmployeeRole {
	return dto.EmployeeRole{
		ID:   domain.ID,
		Role: domain.Role,
	}
}

func ToEmployeeRoleDomain(dto dto.EmployeeRole) domain.EmployeeRole {
	return domain.EmployeeRole{
		ID:   dto.ID,
		Role: dto.Role,
	}
}
