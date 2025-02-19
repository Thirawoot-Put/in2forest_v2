package mapper

import (
	"thirawoot/in2forest_shop_backend/internal/domain"
	"thirawoot/in2forest_shop_backend/internal/dto"
	"thirawoot/in2forest_shop_backend/internal/infras/database/model"
)

func ToEmployeeModel(d domain.Employee) model.Employee {
	return model.Employee{
		Email:          d.Email,
		Password:       d.Password,
		Name:           d.Name,
		Mobile:         d.Mobile,
		EmployeeRoleID: d.Role.ID,
	}
}

func ToEmployeeDomain(m model.Employee) domain.Employee {
	return domain.Employee{
		ID:       m.ID,
		Email:    m.Email,
		Password: m.Password,
		Name:     m.Name,
		Mobile:   m.Mobile,
		Role: domain.EmployeeRole{
			ID:   m.EmployeeRole.ID,
			Name: m.EmployeeRole.Name,
		},
	}
}

func DtoToEmployeeDomain(t dto.Employee) domain.Employee {
	return domain.Employee{
		ID:       t.ID,
		Email:    t.Email,
		Password: t.Password,
		Name:     t.Name,
		Mobile:   t.Mobile,
		Role: domain.EmployeeRole{
			ID:   t.Role.ID,
			Name: t.Role.Name,
		},
	}
}

func DaminToEmployeeDto(d domain.Employee) dto.Employee {
	return dto.Employee{
		ID:       d.ID,
		Email:    d.Email,
		Password: d.Password,
		Name:     d.Name,
		Mobile:   d.Mobile,
		Role: dto.EmployeeRole{
			ID:   d.Role.ID,
			Name: d.Role.Name,
		},
	}
}
