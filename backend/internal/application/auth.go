package application

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	authjwt "thirawoot/in2forest_shop_backend/internal/utils/auth_jwt"
)

type AuthEmployeeAppImpl struct {
	empApp     portin.EmployeeApp
	empRoleApp portin.EmployeeRoleApp
}

func NewAuthEmployeeApp(
	empApp portin.EmployeeApp,
	empRoleApp portin.EmployeeRoleApp,
) portin.AuthEmployeeApp {
	return &AuthEmployeeAppImpl{
		empApp:     empApp,
		empRoleApp: empRoleApp,
	}
}

func (a *AuthEmployeeAppImpl) RegisterAdmin(data dto.Employee) (*dto.AuthResponse, error) {
	role, err := a.empRoleApp.FindByRole("ADMIN")
	if err != nil {
		return nil, err
	}

	data.EmployeeRoleID = role.ID

	emp, err := a.empApp.Create(data)
	if err != nil {
		return nil, err
	}

	token, err := authjwt.SignToken(emp.ID, role.Role)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{AccessToken: token}, nil
}
