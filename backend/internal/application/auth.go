package application

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	authjwt "thirawoot/in2forest_shop_backend/pkg/jwt_util"
	"time"
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
	exp := time.Now().Add(time.Duration(24 * time.Hour))

	token, err := authjwt.GenerateToken(emp.ID, role.Role, exp)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{AccessToken: token}, nil
}
