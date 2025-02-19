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

func ToEmployeeDto(domain domain.Employee) dto.Employee {
	return dto.Employee{
		ID:       domain.ID,
		Email:    domain.Email,
		Password: domain.Password,
		Name:     domain.Name,
		Mobile:   domain.Mobile,

		EmployeeRoleID: domain.EmployeeRoleID,
	}
}
